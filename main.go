package main

import (
	"fmt"

	"example.com/calculator/file_manager"
	"example.com/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChan := make([]chan bool, len(taxRates))

	for idx, taxRate := range taxRates {
		doneChan[idx] = make(chan bool)

		fm := file_manager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdM := cmd_manager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go priceJob.Process(doneChan[idx])

	}

	for _, cha := range doneChan {
		<-cha
	}
	// for range doneChan {

	// }

}
