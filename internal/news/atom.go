package news

import (
	"encoding/xml"
	"informado/internal/pkg/slack"

	log "github.com/sirupsen/logrus"
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

func (a Atom) Print(lastTimeInformadoWasRun int64, sc slack.Channel) error {
	for i := 0; i < len(a.Entry); i++ {
		updated := a.Entry[i].Updated
		updatedInt64, err := dateToEpoch(updated)
		if err != nil {
			return err
		}

		if updatedInt64 > lastTimeInformadoWasRun {
			msg := updated + " " + a.Entry[i].Title.Name + " " + a.Entry[i].Link.Href
			log.Info(msg)
			if sc.ID != "" && sc.Token != "" {
				sc.Msg = msg
				if err := sc.Send(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
