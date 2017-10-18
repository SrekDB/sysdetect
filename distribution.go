package sysdetect

import (
	"fmt"
	"io"
	"io/ioutil"
)

type Distribution string

const (
	DistributionDebian  Distribution = "Debian"
	DistributionFreeBSD              = "FreeBSD"
	DistributionFreeNAS              = "FreeNAS"
	DistributionMacOSX               = "Mac OS X"
)

const (
	DistributionFileOSRelease           = "/etc/os-release"     // https://www.freedesktop.org/software/systemd/man/os-release.html
	DistributionFileOSReleaseAlt        = "/usr/lib/os-release" // https://www.freedesktop.org/software/systemd/man/os-release.html
	DistributionFileDarwinSystemVersion = "/System/Library/CoreServices/SystemVersion.plist"
	DistributionFileDarwinServerVersion = "/System/Library/CoreServices/ServerVersion.plist"
)

func (od Distribution) String() string {
	return string(od)
}

func DistributionFileParse(filename string, r io.Reader) {
	switch filename {
	case DistributionFileOSReleaseAlt:
		fallthrough
	case DistributionFileOSRelease:
		fmt.Println("parse", filename)
		c, _ := ioutil.ReadAll(r)
		fmt.Println(string(c))
		break
	case DistributionFileDarwinSystemVersion:
		break
	case DistributionFileDarwinServerVersion:
		break
	}
}
