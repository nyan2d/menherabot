package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/nyan2d/menherabot/app"
	"github.com/nyan2d/menherabot/config"
	"github.com/nyan2d/menherabot/lol"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cfgpath := "config.json"

	flag.Parse()
	flag.StringVar(&cfgpath, "c", "config.json", "config file path")

	c, err := config.ReadConfigFromFile(cfgpath)
	if err != nil {
		panic(err)
	}

	lol.Init(c.LeagueToken)

	a := app.NewApp(c)
	a.Run()
}
