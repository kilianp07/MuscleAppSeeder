package env

import (
	"fmt"
	"os"
)

var varsNames = []string{
	"YOUTUBE_API_KEY",
}

type Env struct {
	YouTubeApiKey string
}

func New() (*Env, error) {
	e := &Env{}
	if err := e.checkEnv(); err != nil {
		return nil, err
	}

	e.getEnv()

	return e, nil
}
func (e *Env) checkEnv() error {
	for _, v := range varsNames {
		if _, ok := os.LookupEnv(v); !ok {
			return fmt.Errorf("missing env variable %s", v)
		}
	}
	return nil
}

func (e *Env) getEnv() {
	e.YouTubeApiKey = os.Getenv("YOUTUBE_API_KEY")
}
