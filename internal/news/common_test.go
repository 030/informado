package news

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateToEpoch(t *testing.T) {
	inputOutput := map[string]int64{
		"2020-07-30T12:53:50.895-04:00":   1596128030,
		"2020-08-08T00:10:00Z":            1596845400,
		"Fri, 26 Jun 2020 14:00:00 +0000": 1593180000,
		"Sat, 08 Aug 2020 12:06:44 +0200": 1596881204,
		"Wed, 05 Aug 2020 09:09:00 GMT":   1596618540,
	}

	for date, exp := range inputOutput {
		act, _ := dateToEpoch(date)
		assert.Equalf(t, exp, act, date)
	}

	// unhappy
	_, err := dateToEpoch("fffffffffffffffffffff")
	assert.EqualError(t, err, "'fffffffffffffffffffff' cannot be parsed. Check whether the date matches the regex")
}
