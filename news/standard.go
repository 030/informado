package news

import (
	"encoding/xml"
	"fmt"
	"regexp"
)

type Item struct {
	XMLName xml.Name `xml:"item"`
	PubDate string   `xml:"pubDate"`
	Title
	Link string `xml:"link"`
}
type Standard struct {
	XMLName xml.Name `xml:"rss"`
	Item    []Item   `xml:"channel>item"`
}

func (s Standard) Parse(b []byte) (RSS, error) {
	if err := xml.Unmarshal(b, &s); err != nil {
		return Standard{}, err
	}
	return s, nil
}

func (s Standard) Print(date string) error {
	for i := 0; i < len(s.Item); i++ {
		pubDate := s.Item[i].PubDate
		match, err := regexp.MatchString(date, pubDate)
		if err != nil {
			return err
		}
		if match {
			fmt.Println(pubDate + " " + s.Item[i].Title.Name + " " + s.Item[i].Link)
		}
	}
	return nil
}
