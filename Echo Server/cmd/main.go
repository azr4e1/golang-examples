package main

import (
	"echo"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 || args[1] == "network" {
		echo.NetworkEcho()
	} else {
		echo.LocalEcho()
	}
}
