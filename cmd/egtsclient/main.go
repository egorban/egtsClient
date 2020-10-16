package main

import (
	"flag"

	"github.com/egorban/ndtpClient/pkg/egtsclient"
)

func main() {
	var serverAddress string
	var numOIDs int
	var time int
	flag.StringVar(&serverAddress, "s", "localhost:9001", "server address (e.g. 'localhost:9000')")
	flag.IntVar(&numOIDs, "i", 3, "number IDs (e.g. '1')")
	flag.IntVar(&time, "t", 20, "time period seconds (e.g. '20')")
	flag.Parse()
	egtsclient.Start(serverAddress, numOIDs, time)
}
