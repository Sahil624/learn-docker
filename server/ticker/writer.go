package ticker

import "github.com/Sahil624/learn-docker/coingeco"

func readChannel() {
	for price := range coingeco.GetChannel() {
		for k := range price {
			for _, conn := range symbolConnection[k] {
				conn.WriteJSON(price)
			}
		}
	}
}
