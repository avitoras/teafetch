package main

import (
	"fetch/gofetch"
	"fmt"
	"os"
	"os/user"
	"runtime"
	// "strings"
	"github.com/shirou/gopsutil/v4/host"
)

func main() {
	// user@hostname
	username, _ := user.Current()
	hostname, _ := os.Hostname()
	userhost := username.Username + "@" + hostname
	gofetch.PrintWithUnderline(userhost)

	// OS
	dist, _, version, _ := host.PlatformInformation()
	fmt.Println("OS:", dist, version)

	// Kernel
	// kernel, _ := runtime.GOOS
	kernelver, _ := host.KernelVersion()
	fmt.Println("Kernel:", runtime.GOOS,kernelver)
}
