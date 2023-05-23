package main

import (
	"github.com/kilianp07/MuscleAppSeeder/pkg/api"
	csvReader "github.com/kilianp07/MuscleAppSeeder/pkg/csv"
)

func main() {
	data, err := csvReader.GetData()
	if err != nil {
		panic(err)
	}

	if err := api.StartApi(data); err != nil {
		panic(err)
	}

}
