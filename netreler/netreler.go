package netreler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/sparrc/go-ping"
)

const (
	conditionGood           = "good"
	conditionGoodPercentage = 72
	conditionSoSo           = "so-so"
	conditionSoSoPercentage = 49
	conditionPoor           = "poor"
)

const (
	defaultIterationPerSinglePingTest = 200
	defaultIterationPerTest           = 5
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

type TestResult struct {
	TestedHosts []string
	Score       float64
	Meta        AnalyzedResults
	Spent       time.Duration
	Condition   string
}

func (t *TestResult) JSON() ([]byte, error) {
	return json.Marshal(t)
}

func (t *TestResult) PrettyPrintMeta() ([]byte, error) {
	return json.MarshalIndent(t.Meta, "", "\t")
}

func Test(c chan os.Signal) *TestResult {
	mes := time.Now()

	go func() {
		var progress = "="
		for {
			fmt.Printf("%s\r", progress)
			time.Sleep(500 * time.Millisecond)
			progress += "="
		}
	}()

	var ar = make(AnalyzedResults)
	for _, host := range trustedHosts {
		ar[host] = []AnalyzedResult{}
		for i := 0; i < defaultIterationPerTest; i++ {
			pr, err := SinglePingTest(host, c)
			if err != nil {
				log.Print(err)
				continue
			}
			ar[host] = append(ar[host], pr.Analyze())
		}
	}

	score := ar.Analyze()
	condition := conditionPoor
	if score > conditionGoodPercentage {
		condition = conditionGood
	} else if score > conditionSoSoPercentage {
		condition = conditionSoSo
	}

	return &TestResult{
		TestedHosts: trustedHosts,
		Score:       score,
		Meta:        ar,
		Spent:       time.Since(mes),
		Condition:   condition,
	}
}

func SinglePingTest(host string, c chan os.Signal) (*PingResult, error) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return nil, err
	}
	pinger.Interval = time.Millisecond * 10
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
		max := 50
		pinger.Size = rand.Intn(max-min+1) + min
		if counter > defaultIterationPerSinglePingTest {
			pinger.Stop()
			return
		} else {
			counter++
		}
		if pkt != nil {
			//fmt.Println("test")
			pr.Packets = append(pr.Packets, *pkt)
		}
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		pr.Stats = stats
	}
	pinger.Run()

	return pr, nil
}
