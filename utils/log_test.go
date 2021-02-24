package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name        string
	lineString  string
	shouldError bool
}

var cases = []Case{
	{
		name:        "EMPTY STRING",
		lineString:  "",
		shouldError: true,
	},
	{
		name:        "INVALID DATE FORMAT",
		lineString:  "::1 - - [23-Feb-2021:08:09:59 +0800] \"GET / HTTP/1.1\" 200 45",
		shouldError: true,
	},
	{
		name:        "VALID DATE FORMAT",
		lineString:  "::1 - - [23/Feb/2021:08:09:59 +0800] \"GET / HTTP/1.1\" 200 45",
		shouldError: false,
	},
}

func TestLogTimeDiff(t *testing.T) {
	for _, _case := range cases {

		var result bool

		t.Logf(`Testing case %v`, _case.name)
		timeDiff, err := LogTimeDiff(_case.lineString)
		if _case.shouldError {
			t.Logf("expecting error. err = %+v", err)
			result = assert.NotNil(t, err, "expecting error")
		} else {
			t.Logf("expecting no error. err = %+v", err)
			result = assert.Equal(t, nil, err, "expecting no error")
		}

		if !result {
			t.Logf("got result: %+v", result)
			t.Fail()
		}

		t.Logf(`time diff : %+v`, timeDiff)
	}
}
