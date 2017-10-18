package sysdetect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetect(t *testing.T) {
	tsd := &TestSysDetector{}
	assert.Nil(t, Detect(tsd))
}
