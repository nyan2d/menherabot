package main

import (
	"math/rand"
	"time"

	"github.com/nyan2d/menherabot/app"
	"github.com/nyan2d/menherabot/config"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c, err := config.ReadConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}

	a := app.NewApp(c)
	a.Run()
}
