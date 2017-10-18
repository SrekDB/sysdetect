package main

import (
	"github.com/xor-gate/sysdetect"
)

func main() {
	sysdetect.NewLocalSysDetector().Detect()
}
