package app

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	tg "github.com/tucnak/telebot"
	"github.com/wesovilabs/koazee"
)

func (a *App) champikiCommand(m *tg.Message) {
	freeRotate, _ := a.leagueClient.Riot.Champion.GetFreeRotation()
	allChamps, _ := a.leagueClient.DataDragon.GetChampions()

	champNames := []string{}
	for _, champion := range allChamps {
		for _, rotateId := range freeRotate.FreeChampionIDs {
			if strconv.Itoa(rotateId) == champion.Key {
				champNames = append(champNames, champion.Name)
			}
		}
	}

	counter := 0
	xd := koazee.StreamOf(champNames).
		Sort(
			func(a, b string) int {
				return strings.Compare(a, b)
			},
		).
		Map(func(a string) string {
			counter++
			return fmt.Sprintf("%v) %v", counter, a)
		}).
		Out().Val().([]string)


	a.bot.Reply(m, strings.Join(xd, "\n"))
}

func (a *App) rollCommand(m *tg.Message) {
	num := rand.Intn(99) + 1
	a.bot.Reply(m, fmt.Sprintf("%v", num))
}

func (a *App) vacmanCommand(m *tg.Message) {
	vacrate := m.Sender.ID % 100
	a.bot.Reply(m, fmt.Sprintf("Ты вакмен на %v%%", vacrate))
}
