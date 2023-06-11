package main

import (
	"flag"
	"github.com/manuellysuzik/hotel.git/cmd"
)

const (
	startApplication = "guests"
	other            = "rooms"
)

var command string

func main() {

	flag.StringVar(&command, "exec", startApplication, startApplication)

	flag.Parse()

	switch command {
	case startApplication:
		cmd.BootstrapGuestApi()
	case other:
		cmd.BootstrapRoomsApi()
	}
}
