package main

import (
	"github.com/PumpkinSeed/netrel/netreler"
	"os"
	"os/signal"
)

func main() {
	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	pr, _ := netreler.SinglePing("google.com", c)
	pr.Analyze()
}
