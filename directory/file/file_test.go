package file

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

type Case struct {
	name        string
	path     string
	fileBuffer io.Reader
	shouldError bool
	filePrefix string
}

var lineParserCase = []Case{
	{
		name:        "ONE LINER",
		path:     "",
		fileBuffer : strings.NewReader("This is a test string"),
		shouldError: false,
	},
	{
		name:        "MULTI LINE",
		path:     "daodoq.kda!@kd",
		fileBuffer : strings.NewReader("This is a test string \n This is a second test string"),
		shouldError: false,
	},
	{
		name:        "NO FILE",
		path:     "/usr/local/var/log/httpd",
		shouldError: true,
	},

	{
		name:        "EMPTY FILE",
		path:     "/usr/local/var/log/httpd",
		fileBuffer: strings.NewReader(""),
		shouldError: false,
	},
}

var checkIfTargetFileCase = []Case{
	{
		name:        "PATH LENGTH ZERO",
		path:     "",
		filePrefix: "access_log",
		shouldError: true,
	},
	{
		name:        "PREFIX LENGTH ZERO",
		path:     "/usr/local/var/log/httpd",
		filePrefix: "",
		shouldError: true,
	},
	{
		name:        "NOT A TARGET FILE",
		path:     "/usr/local/var/log/httpd/access_log.1613967000",
		filePrefix: "error_log",
		shouldError: false,
	},
	{
		name:        "TARGET FILE",
		path:     "/usr/local/var/log/httpd/access_log.1613967000",
		filePrefix: "access_log",
		shouldError: false,
	},
}

func TestLineParser(t *testing.T) {

	var result bool

	for _, pbc := range lineParserCase {
		t.Logf(`Testing Scenario %s`, pbc.name)
		line,err := LineParser(pbc.fileBuffer)
		if pbc.shouldError {
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

		t.Logf(`parsed line : %+v`, line)
	}
}

func TestCheckIfTargetFile(t *testing.T) {

	var result bool

	for _, _case := range checkIfTargetFileCase {
		t.Logf(`Testing Scenario %s`, _case.name)
		isTarget,err := CheckIfTargetFile(_case.path,_case.filePrefix)
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

		t.Logf(`is target file %v`,isTarget)
	}
}