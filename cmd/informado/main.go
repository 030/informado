package main

import (
	"errors"
	"flag"
	"fmt"
	"informado/internal/news"
	"path/filepath"
	"strings"

	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"

	"github.com/gocarina/gocsv"
)

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

func rss(r news.RSS, url string, lastTimeInformadoWasRun int64) error {
	byte, err := readURL(url)
	if err != nil {
		return err
	}
	a, err := r.Parse(byte)
	if err != nil {
		return err
	}
	if err := a.Print(lastTimeInformadoWasRun); err != nil {
		return err
	}
	return nil
}

type RSSFeeds struct {
	Type  string `csv:"type"`
	URL   string `csv:"url"`
	Error error
}

func csv(f string) ([]*RSSFeeds, error) {
	u, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer u.Close()

	r := []*RSSFeeds{}
	if err := gocsv.UnmarshalFile(u, &r); err != nil {
		return nil, err
	}

	return r, err
}

var wg sync.WaitGroup

func newsItems(c chan *RSSFeeds, r *RSSFeeds, lastTimeInformadoWasRun int64) {
	defer wg.Done()
	var e error

	switch t := r.Type; t {
	case "atom":
		if err := rss(news.Atom{}, r.URL, lastTimeInformadoWasRun); err != nil {
			e = err
		}
	case "standard":
		if err := rss(news.Standard{}, r.URL, lastTimeInformadoWasRun); err != nil {
			e = err
		}
	default:
		e = fmt.Errorf("unsupported type '%v'", t)
	}

	c <- &RSSFeeds{r.Type, r.URL, e}
}

func read(urls []*RSSFeeds, lastTimeInformadoWasRun int64) error {
	c := make(chan *RSSFeeds, len(urls))
	for _, a := range urls {
		wg.Add(1)
		go newsItems(c, a, lastTimeInformadoWasRun)
	}
	wg.Wait()
	close(c)
	for item := range c {
		if item.Error != nil {
			return fmt.Errorf("type: '%v'. URL: '%v', Err: '%v'", item.Type, item.URL, item.Error)
		}
	}
	return nil
}

func parse(home string, lastTimeInformadoWasRun int64) error {
	f := filepath.Join(home, "rss-feed-urls.csv")
	urls, err := csv(f)
	if err != nil {
		return err
	}
	if err := read(urls, lastTimeInformadoWasRun); err != nil {
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

func lastRunTimeAndWriteCurrentTimeToDisk(home string) error {
	f := filepath.Join(home, "last-run-time.txt")

	t, err := lastRun(f)
	if err != nil {
		return err
	}
	log.Infof("informado last run: '%s'", time.Unix(t, 0).Format(time.RFC3339))

	if err := parse(home, t); err != nil {
		return err
	}

	if err := currentTimeToDisk(f); err != nil {
		return err
	}
	return nil
}

func main() {
	home, err := informadoHome()
	if err != nil {
		log.Fatal(err)
	}

	if err := lastRunTimeAndWriteCurrentTimeToDisk(home); err != nil {
		log.Fatal(err)
	}
}
