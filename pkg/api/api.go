package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/schollz/progressbar/v3"
)

type API struct {
	apiUrl string
	user   User
	token  Token
	data   []float64
}

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func StartApi(data []float64) error {
	api := API{
		apiUrl: "http://localhost:8080",
		data:   data,
	}
	api.createUser()
	if err := api.postUser(); err != nil {
		return err
	}

	if err := api.Login(); err != nil {
		return err
	}

	if err := api.postData(); err != nil {
		return err
	}

	fmt.Println("You can now login with the following credentials:")
	fmt.Println("email: ", api.user.Email)
	fmt.Println("Password: ", api.user.Password)

	return nil
}

func (api *API) createUser() {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	api.user = User{
		Name:     nameGenerator.Generate(),
		Surname:  nameGenerator.Generate(),
		Username: nameGenerator.Generate(),
		Password: "test",
		Email:    nameGenerator.Generate() + "@test.com",
	}
}

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

func (api *API) Login() error {
	var posturl = api.apiUrl + "/auth/login"

	data := url.Values{}
	data.Set("email", api.user.Email)
	data.Set("password", api.user.Password)

	r, err := http.PostForm(posturl, data)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, &api.token); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	return nil
}

func (api *API) postData() error {
	var posturl = api.apiUrl + "/weight"

	startDate := time.Now().Add(-24 * time.Hour * (time.Duration(len(api.data) / 2)))

	fmt.Println("Posting weight data to API")
	bar := progressbar.Default(int64(len(api.data)))

	for _, data := range api.data {
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
