package sysdetect

type SysInfo interface {
	String() string
	Distribution() Distribution
}
