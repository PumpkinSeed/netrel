package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/PumpkinSeed/netrel/netreler"
)

func main() {
	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	result := netreler.Test(c)
	_, err := result.JSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("")
	fmt.Println(result.Score)
	fmt.Println(result.Spent)
	//fmt.Println(string(jso))
}
