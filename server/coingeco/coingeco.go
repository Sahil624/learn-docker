package coingeco

func UpdateCoins() {
	if _, err := getListOfCoins(); err != nil {
		return
	}

}

func fetchPrices() {
	//cgClient.SimplePrice()
}
