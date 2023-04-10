package public

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Price struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func GetBinancePrice(currency string) (float64, error) {
	var url = "https://api.binance.com/api/v3/ticker/price?symbol=%s"

	switch currency {
	case "ETH":
		url = fmt.Sprintf(url, "ETHUSDT")
	case "BNB":
		url = fmt.Sprintf(url, "BNBUSDT")
	}

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var price Price
	err = json.NewDecoder(resp.Body).Decode(&price)
	if err != nil {
		return 0, err
	}

	priceFloat, err := strconv.ParseFloat(price.Price, 64)
	if err != nil {
		return 0, err
	}

	return priceFloat, nil
}
