package prices

import (
	"fmt"

	"example.com/calculator/conversion"
	"example.com/calculator/io_manager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64              `json:"tax_rate"`
	InputPrices       []float64            `json:"input_prices"`
	TaxIncludedPrices map[string]string    `json:"tax_included_prices"`
	IOManger          io_manager.IOManager `json:"-"` // exclude this key and value
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManger.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func NewTaxIncludedPriceJob(iom io_manager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManger:    iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	err := job.IOManger.WriteResult(job)

	if err != nil {
		errorChan <- err
		return
	}

	doneChan <- true
}
