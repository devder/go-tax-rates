package main

import (
	"fmt"

	"example.com/calculator/file_manager"
	"example.com/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChan := make([]chan bool, len(taxRates))
	errorChan := make([]chan error, len(taxRates))

	for idx, taxRate := range taxRates {
		doneChan[idx] = make(chan bool)
		errorChan[idx] = make(chan error)

		fm := file_manager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdM := cmd_manager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go priceJob.Process(doneChan[idx], errorChan[idx])

	}

	// to handle a situation where there is more than one channel
	// the case that gives a result earlier, wins and the other case is ignored
	for i := range taxRates {
		select {
		case err := <-errorChan[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChan[i]:
			fmt.Println("done")
		}
	}

}
