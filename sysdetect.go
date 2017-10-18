package sysdetect

import (
	"bytes"
)

type SysDetector interface {
	FileExists(filename string, allowEmpty bool) bool
	ReadFile(filename string) ([]byte, error)
}

func Detect(sd SysDetector) SysInfo {
	if sd.FileExists(DistributionFileOSRelease, false) {
		osrelease, err := sd.ReadFile(DistributionFileOSRelease)
		if err == nil {
			DistributionFileParse(DistributionFileOSRelease, bytes.NewReader(osrelease))
		}
	}

	return nil
}
