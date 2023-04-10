package public

import (
	"fmt"
	"testing"
)

func TestHttpRequest(t *testing.T) {
	// ETH
	ethPrice, err := GetBinancePrice("ETH")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("ETH/USDT price: %.2f\n", ethPrice)

	// BNB
	bnbPrice, err := GetBinancePrice("BNB")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("BNB/USDT price: %.2f\n", bnbPrice)
}
