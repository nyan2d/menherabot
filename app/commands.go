package app

import (
	"fmt"
	tg "github.com/tucnak/telebot"
	"math/rand"
	"strconv"
)

func (a *App) champikiCommand(m *tg.Message) {
	freeRotate, _ := a.leagueClient.Riot.Champion.GetFreeRotation()
	allChamps, _ := a.leagueClient.DataDragon.GetChampions()

	res := ""
	for _, champion := range allChamps {
		for _, rotateId := range freeRotate.FreeChampionIDs {
			if strconv.Itoa(rotateId) == champion.Key {
				res += champion.Name + "\n"
			}
		}
	}
	a.bot.Reply(m, res)
}

func (a *App) rollCommand(m *tg.Message) {
	num := rand.Intn(99) + 1
	a.bot.Reply(m, fmt.Sprintf("%v", num))
}

func (a *App) vacmanCommand(m *tg.Message) {
	vacrate := m.Sender.ID % 100
	a.bot.Reply(m, fmt.Sprintf("Ты вакмен на %v%%", vacrate))
}
