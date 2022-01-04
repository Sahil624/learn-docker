package ticker

import "github.com/Sahil624/learn-docker/coingeco"

func readChannel() {
	for price := range coingeco.GetChannel() {
		for k := range price {
			for _, conn := range symbolConnection[k] {
				if conn != nil && conn.Conn != nil {
					conn.WriteJSON(price)
				}
			}
		}
	}
}
