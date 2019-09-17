package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/PumpkinSeed/netrel/netreler"
)

var (
	flagPrintMeta bool
)

func init() {
	flag.BoolVar(&flagPrintMeta, "print-meta", false, "Printing meta results")
	flag.Parse()
}

func main() {
	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	result := netreler.Test(c)

	if flagPrintMeta {
		meta, err := result.PrettyPrintMeta()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(meta))
	}
	fmt.Printf("\n\n")
	fmt.Printf("Tested hosts:\n")
	for _, host := range result.TestedHosts {
		fmt.Printf("\t%s\n", host)
	}
	fmt.Printf("Final score of internet reliability: %f%% \n", result.Score)
	fmt.Printf("Test spent: %s\n", result.Spent)
}
