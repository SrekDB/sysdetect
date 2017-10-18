package main

import (
	"fmt"
	"github.com/xor-gate/sysdetect"
)

func main() {
	fmt.Println(sysdetect.NewLocalSysDetector().Detect())
}
