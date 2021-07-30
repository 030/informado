package news

import (
	"fmt"
	"regexp"
	"time"
)

type Title struct {
	Name string `xml:"title"`
}

func dateToEpoch(date string) (int64, error) {
	var t time.Time
	var err error

	dateZ := regexp.MustCompile(`^.*\dT\d.*$`)
	dateZone := regexp.MustCompile(`^.*GMT$`)
	dateNumericZone := regexp.MustCompile(`^.*(\-|\+)\d+$`)

	switch {
	case dateZ.MatchString(date):
		t, err = time.Parse(time.RFC3339, date)
	case dateZone.MatchString(date):
		t, err = time.Parse(time.RFC1123, date)
	case dateNumericZone.MatchString(date):
		t, err = time.Parse(time.RFC1123Z, date)
	default:
		return 0, fmt.Errorf("'%v' cannot be parsed. Check whether the date matches the regex", date)
	}

	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
