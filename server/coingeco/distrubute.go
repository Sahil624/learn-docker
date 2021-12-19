package coingeco

var symbolChan = make(chan map[string]float32, 50)

func GetChannel() chan map[string]float32 {
	return symbolChan
}
