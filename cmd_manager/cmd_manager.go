package cmd_manager

import "fmt"

type CMDManager struct{}

func New() *CMDManager {
	return &CMDManager{}
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices\nConfirm every price with ENTER\nEnter 0 when complete")

	var prices []string

	for {
		var price string
		fmt.Println("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}
