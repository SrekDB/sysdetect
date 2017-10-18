package sysdetect

import (
	"bytes"
	"fmt"
)

type SysDetector interface {
	FileExists(filename string, allowEmpty bool) bool
	ReadFile(filename string) ([]byte, error)
	Sysname() Sysname
	RunCommand(name string, arg ...string) (string, error)
}

func Detect(sd SysDetector) SysInfo {
	sysname := sd.Sysname()

	switch sysname {
	case SysnameDarwin:
		fmt.Println("sysname:", sysname)
		systemversion, err := sd.ReadFile(DistributionFileDarwinSystemVersion)
		if err == nil {
			fmt.Println("found:\n", string(systemversion))
		}
	case SysnameLinux:
		osrelease, err := sd.ReadFile(DistributionFileOSRelease)
		if err == nil {
			DistributionFileParse(DistributionFileOSRelease, bytes.NewReader(osrelease))
		}
	case SysnameUnknown:
		fmt.Println("sysname:", sysname)
	default:
	}

	return nil
}
