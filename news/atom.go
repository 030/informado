package news

import (
	"encoding/xml"
	"fmt"
	"regexp"
)

type Link struct {
	Href string `xml:"href,attr"`
}
type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Updated string   `xml:"updated"`
	Title
	Link Link `xml:"link"`
}
type Atom struct {
	Feed  xml.Name `xml:"feed"`
	Entry []Entry  `xml:"entry"`
}

func (a Atom) Parse(b []byte) (RSS, error) {
	if err := xml.Unmarshal(b, &a); err != nil {
		return Atom{}, err
	}
	return a, nil
}

func (a Atom) Print(date string) error {
	for i := 0; i < len(a.Entry); i++ {
		updated := a.Entry[i].Updated
		match, err := regexp.MatchString(date, updated)
		if err != nil {
			return err
		}
		if match {
			fmt.Println(updated + " " + a.Entry[i].Title.Name + " " + a.Entry[i].Link.Href)
		}
	}
	return nil
}
