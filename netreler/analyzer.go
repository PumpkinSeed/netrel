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

func (pr *PingResult) Analyze() AnalyzedResult {
	var ar = AnalyzedResult{
		Result: pr,
	}

	rttMedina := pr.medianRtt()
	var biggerThanMedianCounter = 0
	for _, packet := range pr.Packets {
		if float64(packet.Rtt) > rttMedina {
			biggerThanMedianCounter++
		}
	}
	var calculated = float64(biggerThanMedianCounter)/float64(len(pr.Packets))
	var calculatedInt = int(calculated*100)
	var	final int
	if calculatedInt > 50 {
		part := 50 - (100 - calculatedInt)
		final = 100 - part*2
	} else {
		part := 50 - calculatedInt
		final = 100 - part*2
	}

	ar.Score = final

	return ar
}

func (pr *PingResult) medianRtt() float64 {
	var sumRtt = 0
	for _, packet := range pr.Packets {
		sumRtt += int(packet.Rtt)
	}

	return float64(sumRtt) / float64(len(pr.Packets))
}

type AnalyzedResults map[string][]AnalyzedResult

func (ar AnalyzedResults) Analyze() float64 {
	var sum int
	var length = 0
	for _, arAction := range ar {
		length += len(arAction)
		for _, singleAction := range arAction {
			sum += singleAction.Score
		}
	}
	return float64(sum) / float64(length)
}

