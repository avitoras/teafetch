package main

import (
	"teafetch/pwul"
	"teafetch/utils"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"github.com/shirou/gopsutil/v4/host"
)

func main() {
	// user@hostname
	username, _ := user.Current()
	hostname, _ := os.Hostname()
	userhost := username.Username + "@" + hostname
	PWUL.PrintWithUnderline(userhost)

	// OS
	dist, _, version, _ := host.PlatformInformation()
	fmt.Println("OS:", utils.CapitalizeFirstLetter(dist), version)

	// Kernel
	kernelver, _ := host.KernelVersion()
	fmt.Println("Kernel:", utils.CapitalizeFirstLetter(runtime.GOOS),kernelver)
}
