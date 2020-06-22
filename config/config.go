package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type Config struct {
	BotToken    string `json:"bot_token"`
	LeagueToken string `json:"league_token"`
}

func ReadConfig(r io.Reader) (*Config, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	c := Config{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func ReadConfigFromFile(file string) (*Config, error) {
	ffile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer ffile.Close()

	return ReadConfig(ffile)
}
