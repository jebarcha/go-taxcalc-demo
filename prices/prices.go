package prices

import (
	"fmt"
	"price-calc/conversion"
	"price-calc/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	//fmt.Println(result)
	job.TaxIncludedPrices = result

	filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
