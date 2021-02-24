package directory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const name = `./test`

func TestDirectoryWalk(t *testing.T) {

	var result bool

	da,err := NewDirectoryAction()
	result = assert.Equal(t, nil, err, "expecting no error")
	if err != nil {
		t.Logf("error new directory action. err = %+v", err)
	}

	err = da.Walk(name)
	result = assert.Equal(t, nil, err, "expecting no error")
	if err != nil {
		t.Logf("error directory walk. err = %+v", err)
	}

	if !result {
		t.Fail()
	}
}
