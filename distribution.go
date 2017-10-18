package sysdetect

type Distribution string

const (
	DistributionDebian  Distribution = "Debian"
	DistributionFreeBSD Distribution = "FreeBSD"
)

func (od Distribution) String() string {
	return string(od)
}
