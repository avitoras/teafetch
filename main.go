package main

import (
	"fetch/gofetch"
	"fmt"
	"os"
	"os/user"
)

func main() {
	// user@hostname
	user, _ := user.Current()
	host, _ := os.Hostname()
	userhost := user.Username + "@" + host
	gofetch.PrintWithUnderline(userhost)

	// OS
	///fmt.Println("OS: ",os.ReadFile())
}
