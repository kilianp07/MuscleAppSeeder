package csvReader

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	exerciseModel "github.com/kilianp07/MuscleApp/models/Exercise"
	randomVideo "github.com/kilianp07/MuscleAppSeeder/pkg/youtube"
)

type Data struct {
	Weight    []float64              `json:"weight"`
	Exercises []exerciseModel.Create `json:"exercises"`
}

func GetData() (*Data, error) {
	d := &Data{}

	if err := d.getWeights(); err != nil {
		return nil, err
	}

	if err := d.getExercises(); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Data) getWeights() error {

	var (
		weights []float64
		data    float64
	)

	csvFile, err := os.Open("./weight.csv")
	if err != nil {
		return err
	}
	fmt.Println("Successfully Opened Weight CSV file")
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

	d.Weight = weights

	return nil
}

func (d *Data) getExercises() error {

	var (
		exercises []exerciseModel.Create
	)

	csvFile, err := os.Open("./exercise.csv")
	if err != nil {
		return err
	}
	fmt.Println("Successfully Opened Exercises CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	videos, err := randomVideo.Get(len(csvLines))
	if err != nil {
		return err
	}

	i := 0
	for _, line := range csvLines {

		exercise := exerciseModel.Create{
			Title:       line[1],
			Description: line[2],
			Video:       videos[i],
			Difficulty:  uint(rand.Intn(10)),
			Member:      line[4],
			Type:        line[3],
		}

		if i == len(videos)-1 {
			i = 0
		} else {
			i++
		}

		exercises = append(exercises, exercise)
	}

	d.Exercises = exercises

	return nil
}
