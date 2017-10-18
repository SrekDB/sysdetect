package sysdetect

import (
	"io/ioutil"
	"os/exec"
	"strings"
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

func (od *LocalSysDetector) Sysname() Sysname {
	output, err := od.RunCommand("uname", "-s")
	if err != nil {
		return SysnameUnknown
	}
	return Sysname(strings.Trim(output, " \r\n"))
}

func (sd *LocalSysDetector) RunCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

var _ SysDetector = &LocalSysDetector{}
