package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/schollz/progressbar/v3"
)

func (api *API) postUser() error {
	var posturl = api.apiUrl + "/user"

	jsonUser, err := json.Marshal(api.user)
	if err != nil {
		return err
	}

	// Post request to create user using api.user
	r, err := http.Post(posturl, "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return nil
}

func (api *API) postWeights() error {
	var posturl = api.apiUrl + "/weight"

	startDate := time.Now().Add(-24 * time.Hour * (time.Duration(len(api.data.Weight) / 2)))

	fmt.Println("Posting weight data to API")
	bar := progressbar.Default(int64(len(api.data.Weight)))

	for _, data := range api.data.Weight {
		startDate = startDate.Add(24 * time.Hour)
		jsonData, err := json.Marshal(map[string]interface{}{
			"value": data,
			"date":  startDate.Unix(),
		})
		if err != nil {
			return err
		}

		req, err := http.NewRequest("POST", posturl, bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}
		req.Header.Add("Authorization", "Bearer "+api.token.Token)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()
		bar.Add(1)
	}
	return nil
}

func (api *API) postExercises() error {
	var posturl = api.apiUrl + "/exercise"
	bar := progressbar.Default(int64(len(api.data.Exercises)))

	fmt.Println("Posting exercise data to API")
	for _, data := range api.data.Exercises {
		// Marshal Exercise struct to json
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}

		req, err := http.NewRequest("POST", posturl, bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}
		req.Header.Add("Authorization", "Bearer "+api.token.Token)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()
		bar.Add(1)
	}
	return nil
}
