package main

import (
	"errors"
	"flag"
	"fmt"
	"informado/internal/news"
	"informado/internal/pkg/slack"
	"path/filepath"
	"strings"
	"sync"

	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/gocarina/gocsv"
)

type RSSFeed struct {
	Type  string `csv:"type"`
	URL   string `csv:"url"`
	Error error
}

type informado struct {
	home         string
	lastRunTime  int64
	rssFeed      *RSSFeed
	slackChannel slack.Channel
}

var errs []error
var wg sync.WaitGroup

func readURL(u string) ([]byte, error) {
	var bodyBytes []byte

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	// Required to parse Reddit feeds
	req.Header.Set("user-agent", "hello:world:v0.0 (by /u/ocelost)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("statusCode not 200, but '%v'", resp.StatusCode)
	}
	return bodyBytes, nil
}

func (i *informado) rss(r news.RSS) error {
	byte, err := readURL(i.rssFeed.URL)
	if err != nil {
		return err
	}
	a, err := r.Parse(byte)
	if err != nil {
		return err
	}
	if err := a.Print(i.lastRunTime, i.slackChannel); err != nil {
		return err
	}
	return nil
}

func csv(f string) ([]*RSSFeed, error) {
	u, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer u.Close()

	r := []*RSSFeed{}
	if err := gocsv.UnmarshalFile(u, &r); err != nil {
		return nil, err
	}

	return r, err
}

func (i *informado) newsItems() {
	switch t := i.rssFeed.Type; t {
	case "atom":
		if err := i.rss(news.Atom{}); err != nil {
			errs = append(errs, err)
		}
	case "standard":
		if err := i.rss(news.Standard{}); err != nil {
			errs = append(errs, err)
		}
	default:
		errs = append(errs, fmt.Errorf("unsupported type '%v'", t))
	}
}

func (i *informado) read(feeds []*RSSFeed) error {
	for _, feed := range feeds {
		wg.Add(1)
		go func(feed *RSSFeed) {
			defer wg.Done()
			i.rssFeed = feed
			i.newsItems()
		}(feed)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *informado) parse() error {
	f := filepath.Join(i.home, "rss-feed-urls.csv")
	rssFeeds, err := csv(f)
	if err != nil {
		return err
	}
	if err := i.read(rssFeeds); err != nil {
		return err
	}
	return nil
}

func currentTimeToDisk(f string) error {
	now := time.Now()
	epoch := now.Unix()

	file, err := os.Create(f)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(strconv.FormatInt(epoch, 10))
	if err != nil {
		return err
	}

	return nil
}

func lastRun(f string) (int64, error) {
	if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(f)
		if err != nil {
			return 0, err
		}
		defer file.Close()

		_, err = file.WriteString("0")
		if err != nil {
			return 0, err
		}
	}

	date, err := ioutil.ReadFile(f)
	if err != nil {
		return 0, err
	}

	n, err := strconv.ParseInt(strings.TrimSuffix(string(date), "\n"), 10, 64)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func init() {
	log.SetReportCaller(true)
}

func informadoHome() (string, error) {
	informadoHome := flag.String("home", "", "the home folder that contains the RSS Feed URLs CSV file")
	flag.Parse()

	if *informadoHome == "" {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}

		*informadoHome = filepath.Join(home, ".informado")
	}

	return *informadoHome, nil
}

func (i *informado) lastRunTimeAndWriteCurrentTimeToDisk(home string) error {
	f := filepath.Join("/tmp", "informado", "last-run-time.txt")

	t, err := lastRun(f)
	if err != nil {
		return err
	}
	log.Infof("informado last run: '%s'", time.Unix(t, 0).Format(time.RFC3339))
	i.lastRunTime = t

	if err := i.parse(); err != nil {
		return err
	}

	if err := currentTimeToDisk(f); err != nil {
		return err
	}
	return nil
}

func slackCreds(home string) (slack.Channel, error) {
	viper.SetConfigName("creds")
	viper.SetConfigType("yml")
	viper.AddConfigPath(home)

	if err := viper.ReadInConfig(); err != nil {
		return slack.Channel{}, err
	}

	return slack.Channel{ID: viper.GetString("slackChannel"), Token: viper.GetString("slackToken")}, nil
}

func main() {
	home, err := informadoHome()
	if err != nil {
		log.Fatal(err)
	}
	s, err := slackCreds(home)
	if err != nil {
		log.Fatal(err)
	}
	i := informado{home: home, slackChannel: s}

	if err := i.lastRunTimeAndWriteCurrentTimeToDisk(home); err != nil {
		log.Fatal(err)
	}
}
