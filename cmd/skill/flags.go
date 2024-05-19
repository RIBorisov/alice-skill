package main

import (
	"flag"
)

var flagRunAddr string
var flagLogLevel string

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&flagLogLevel, "f", "INFO", "logger level")
	flag.Parse()
}
