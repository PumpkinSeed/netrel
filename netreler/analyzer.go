package netreler

const (
	increase = iota
	decrease
)

type AnalyzedResult struct {
	Result *PingResult
	Score int
}

type countedMedian struct {
	typ int
	num int
}

func (pr *PingResult) Analyze() *AnalyzedResult {
	var ar = &AnalyzedResult{
		Result: pr,
	}

	//res, _ := json.Marshal(pr)
	// fmt.Println(pr.medians())
	medianNBytes, _ := pr.medians()
	var cmNBytes []countedMedian
	for _, packet := range pr.Packets {
		//fmt.Println(packet.Nbytes)
		//asd := int((float64(packet.Nbytes)/float64(medianNBytes)) * 100)
		if medianNBytes < packet.Nbytes {
			cmNBytes = append(cmNBytes, countedMedian{
				typ: increase,
				num: int((float64(packet.Nbytes)/float64(medianNBytes)) * 100),
			})
			//asd := int((float64(packet.Nbytes)/float64(medianNBytes)) * 100)
			//fmt.Printf("Increase - NBytes: %d, Median: %d, num: %d\n", packet.Nbytes, medianNBytes, asd)
		} else if medianNBytes > packet.Nbytes {
			cmNBytes = append(cmNBytes, countedMedian{
				typ: decrease,
				num: int((float64(medianNBytes)/float64(packet.Nbytes)) * 100),
			})
			//asd := int((float64(medianNBytes)/float64(packet.Nbytes)) * 100)
			//fmt.Printf("Decrease - NBytes: %d, Median: %d, num: %d\n", packet.Nbytes, medianNBytes, asd)
		}

	}
	// fmt.Println(string(res))

	return ar
}

type AnalyzedResults map[string][]AnalyzedResult

func (ar *AnalyzedResults) Analyze() {

}

func (pr *PingResult) medians() (int, int) {
	var sumNBytes = 0
	var sumRtt = 0
	for _, packet := range pr.Packets {
		sumNBytes += packet.Nbytes
		sumRtt += int(packet.Rtt)
	}

	medianNBytes := sumNBytes / len(pr.Packets)
	medianRtt := sumRtt / len(pr.Packets)
	return medianNBytes, medianRtt
}