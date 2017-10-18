package sysdetect

import (
	"io/ioutil"
)

type LocalSysDetector struct {
}

func NewLocalSysDetector() *LocalSysDetector {
	return &LocalSysDetector{}
}

func (od *LocalSysDetector) FileExists(filename string, allowEmpty bool) bool {
	c, err := od.ReadFile(filename)
	if err != nil {
		return false
	}

	if !allowEmpty && len(c) == 0 {
		return false
	}

	return true
}

func (od *LocalSysDetector) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (od *LocalSysDetector) Detect() SysInfo {
	return Detect(od)
}

var _ SysDetector = &LocalSysDetector{}
