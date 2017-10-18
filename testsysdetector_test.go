package sysdetect

type TestSysDetector struct {
}

func (od *TestSysDetector) AddFile(filepath, content string) {
}

func (od *TestSysDetector) FileExists(filename string, allowEmpty bool) bool {
	return false
}

func (od *TestSysDetector) ReadFile(filename string) ([]byte, error) {
	return nil, nil
}

func (sd *TestSysDetector) Sysname() Sysname {
	return ""
}

func (sd *TestSysDetector) RunCommand(name string, arg ...string) (string, error) {
	return ""
}

var _ SysDetector = &TestSysDetector{}
