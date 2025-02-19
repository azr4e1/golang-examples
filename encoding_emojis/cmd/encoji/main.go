package main

import (
	"github.com/azr4e1/encoji"
	"os"
)

func main() {
	os.Exit(encoji.Main(os.Stdin, os.Stdout, os.Stderr))
}
