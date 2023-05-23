package csvReader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func GetData() ([]float64, error) {

	var (
		weights []float64
		data    float64
	)

	csvFile, err := os.Open("./weight.csv")
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {

		if data, err = strconv.ParseFloat(line[2], 64); err != nil {
			continue
		}

		weights = append(weights, data*0.453592)
	}

	return weights, nil
}
