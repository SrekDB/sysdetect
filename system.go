package sysdetect

type Sysname string

const (
	SysnameUnknown Sysname = "Unknown"
	SysnameDarwin          = "Darwin"
	SysnameLinux           = "Linux"
	SysnameFreeBSD         = "FreeBSD"
	SysnameOpenBSD         = "OpenBSD"
	SysnameNetBSD          = "NetBSD"
)

func (s Sysname) String() string {
	switch s {
	case SysnameDarwin:
		fallthrough
	case SysnameFreeBSD:
		fallthrough
	case SysnameOpenBSD:
		fallthrough
	case SysnameNetBSD:
		fallthrough
	default:
		return string(s)
	}
	return string(SysnameUnknown)
}
