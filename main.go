package main

import (
	"030/informado/news"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

const (
	allDayHours = "[0-2][0-9]"
)

func rss(r news.RSS, url, date string) error {
	err, a := r.Parse(url)
	if err != nil {
		return err
	}
	a.Print(date)
	return nil
}

func hello() string {
	return "world"
}

type RSSFeeds struct {
	Type string `csv:"type"`
	URL  string `csv:"url"`
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

func read(urls []*RSSFeeds, date string) error {
	for _, u := range urls {
		switch t := u.Type; t {
		case "atom":
			if err := rss(news.Atom{}, u.URL, date); err != nil {
				return err
			}
		case "standard":
			if err := rss(news.Standard{}, u.URL, date); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Unsupported type '%v'", t)
		}
	}
	return nil
}

func today() string {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	dayMonthYear := currentTime.Format("02 Jan 2006")
	return "^(" + currentDate + "T" + allDayHours + "|.*" + dayMonthYear + " " + allDayHours + "):.*$"
}

func parse(input, date string) {
	urls, err := csv(input)
	if err != nil {
		log.Fatal(err)
	}
	if err := read(urls, date); err != nil {
		log.Fatal(err)
	}
}

func main() {
	date := flag.String("date", "", "Get the RSS feeds from a certain date: '^(2020-06-26T[0-2][0-9]|.*26 Jun 2020 [0-2][0-9]).*$'")
	input := flag.String("file", "informado.csv", "The file that contains a list of RSS URLs")

	if *date == "" {
		*date = today()
	}

	flag.Parse()

	parse(*input, *date)
}
