package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kilianp07/MuscleAppSeeder/pkg/api"
	csvReader "github.com/kilianp07/MuscleAppSeeder/pkg/csv"
	"github.com/kilianp07/MuscleAppSeeder/pkg/env"
)

func main() {
	// Get env variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, trying to load from system environment variables")
	}
	_, err = env.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := csvReader.GetData()
	if err != nil {
		panic(err)
	}

	if err := api.StartApi(data); err != nil {
		panic(err)
	}
}
