package sysdetect

import (
	"io"
	"io/ioutil"
	"fmt"
)

type Distribution string

const (
	DistributionDebian  Distribution = "Debian"
	DistributionFreeBSD Distribution = "FreeBSD"
	DistributionOSX     Distribution = "OSX"
	DistributionMacOS   Distribution = "macOS"
)

const (
	DistributionFileOSRelease = "/etc/os-release" // https://www.freedesktop.org/software/systemd/man/os-release.html
	DistributionFileDarwinSystemVersion = "/System/Library/CoreServices/SystemVersion.plist"
	DistributionFileDarwinServerVersion = "/System/Library/CoreServices/ServerVersion.plist"
)

func (od Distribution) String() string {
	return string(od)
}

func DistributionFileParse(filename string, r io.Reader) {
	switch filename {
	case DistributionFileOSRelease:
		fmt.Println("parse", DistributionFileOSRelease)
		c, _ := ioutil.ReadAll(r)
		fmt.Println(string(c))
		break
	case DistributionFileDarwinSystemVersion:
		break
	case DistributionFileDarwinServerVersion:
		break
	}
}
