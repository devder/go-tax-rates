package main

import (
	"example.com/calculator/cmd_manager"
	"example.com/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := file_manager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdM := cmd_manager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdM, taxRate)
		priceJob.Process()

	}
}
