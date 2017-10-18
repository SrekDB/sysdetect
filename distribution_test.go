package sysdetect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistribution(t *testing.T) {
	assert.Equal(t, "Debian", DistributionDebian.String())
	assert.Equal(t, "FreeBSD", DistributionFreeBSD.String())
}
