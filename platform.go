package sysdetect

type Platform interface {
	Hostname() string
}
