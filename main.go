package main

import (
	"flag"
	"github.com/manuellysuzik/hotel.git/cmd"
)

const (
	startApplication = "app"
	other            = "other"
)

var command string

func main() {

	flag.StringVar(&command, "c", startApplication, startApplication)

	flag.Parse()

	switch command {
	case startApplication:
		cmd.Start
	case other:
		cmd.ExecuteOther()
	}
}
