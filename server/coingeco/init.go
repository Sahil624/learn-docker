package coingeco

import (
	coingecko "github.com/superoo7/go-gecko/v3"
	"net/http"
	"time"
)

var cgClient *coingecko.Client

func init() {

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	cgClient = coingecko.NewClient(httpClient)

	go startPolling()
}
