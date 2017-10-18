package sysdetect

type SysDetector interface {
	FileExists(filename string, allowEmpty bool) bool
	ReadFile(filename string) ([]byte, error)
}

func Detect(sd SysDetector) SysInfo {
	return nil
}
