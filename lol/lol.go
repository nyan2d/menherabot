package lol

import (
	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
)

var client *golio.Client
var apitoken string

func Init(token string) {
	client = golio.NewClient(token, golio.WithRegion(api.RegionEuropeWest))
}

func GetFreeChampPool() ([]string, error) {
	result := []string{}
	freeChamps, err := client.Riot.Champion.GetFreeRotation()
	if err != nil {
		return result, err
	}

	for _, v := range freeChamps.FreeChampionIDs {
		result = append(result, champions[v])
	}

	return result, nil
}
