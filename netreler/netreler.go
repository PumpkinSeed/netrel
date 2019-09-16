package netreler

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"math/rand"
	"os"
	"time"
)

const (
	defaultIterationPerTest = 10
)

var (
	trustedHosts = []string{
		"1.1.1.1",
		"google-public-dns-a.google.com",
		"google-public-dns-b.google.com",
		"139.130.4.5",
	}
)

type PingResult struct {
	Stats   *ping.Statistics
	Packets []ping.Packet
}

func SinglePing(host string, c chan os.Signal) (*PingResult, error) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return nil, err
	}
	pinger.Interval = time.Millisecond * 100
	pinger.SetPrivileged(true)

	go func() {
		for _ = range c {
			pinger.Stop()
		}
	}()

	var pr = &PingResult{}

	var counter = 0
	pinger.OnRecv = func(pkt *ping.Packet) {
		rand.Seed(time.Now().UnixNano())
		min := 24
		max := 110
		pinger.Size = rand.Intn(max - min + 1) + min
		if counter > defaultIterationPerTest {
			pinger.Stop()
			return
		} else {
			counter++
		}
		if pkt != nil {
			fmt.Println("test")
			pr.Packets = append(pr.Packets, *pkt)
		}
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Println("testfinish")
		pr.Stats = stats
	}
	pinger.Run()

	return pr, nil
}
