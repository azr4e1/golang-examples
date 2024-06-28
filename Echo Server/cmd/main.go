package main

import (
	"echo"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 || args[1] == "network" {
		echo.TCPEcho()
	} else if len(args) <= 1 || args[1] == "stdio" {
		echo.STDIOEcho()
	} else {
		echo.UDSEcho()
	}
}
