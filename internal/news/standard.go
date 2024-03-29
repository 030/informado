package news

import (
	"encoding/xml"
	"informado/internal/pkg/slack"

	log "github.com/sirupsen/logrus"
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

func (s Standard) Print(lastTimeInformadoWasRun int64, sc slack.Channel) error {
	for i := 0; i < len(s.Item); i++ {
		updated := s.Item[i].PubDate
		updatedInt64, err := dateToEpoch(updated)
		if err != nil {
			return err
		}

		if updatedInt64 > lastTimeInformadoWasRun {
			msg := updated + " " + s.Item[i].Title.Name + " " + s.Item[i].Link
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
