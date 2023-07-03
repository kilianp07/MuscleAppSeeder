package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/goombaio/namegenerator"
	csvReader "github.com/kilianp07/MuscleAppSeeder/pkg/csv"
)

type API struct {
	apiUrl string
	user   User
	token  Token
	data   *csvReader.Data
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

func StartApi(data *csvReader.Data) error {
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
	if err := api.postWeights(); err != nil {
		return err
	}

	if err := api.postExercises(); err != nil {
		return err
	}

	return nil
}
