package app

import (
	"fmt"
	"github.com/KnutZuidema/golio/datadragon"
	tg "github.com/tucnak/telebot"
	"github.com/wesovilabs/koazee"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (a *App) rotationCommand(m *tg.Message) {
	freeRotate, err := a.leagueClient.Riot.Champion.GetFreeRotation()
	if err != nil {
		a.bot.Reply(m, err)
	}
	allChamps, err := a.leagueClient.DataDragon.GetChampions()
	if err != nil {
		a.bot.Reply(m, err)
	}

	counter := 0
	rotation := koazee.StreamOf(freeRotate.FreeChampionIDs).
		Map(func(a int) datadragon.ChampionData {
			for _, champ := range allChamps {
				if champ.Key == strconv.Itoa(a) {
					return champ
				}
			}
			return datadragon.ChampionData{}
		}).
		Sort(func (a, b datadragon.ChampionData) int {
			return strings.Compare(a.Name, b.Name)
		}).
		Map(func (data datadragon.ChampionData) string {
			counter++
			return fmt.Sprintf("%v) %v", counter, data.Name)
		}).Out().Val().([]string)
	a.bot.Reply(m, strings.Join(rotation, "\n"))
}

func (a *App) rollCommand(m *tg.Message) {
	num := rand.Intn(99) + 1
	a.bot.Reply(m, fmt.Sprintf("%v", num))
}

func (a *App) vacmanCommand(m *tg.Message) {
	t := time.Now()
	vacrate := (t.Year()+int(t.Month())+t.Day() * m.Sender.ID) % 100
	a.bot.Reply(m, fmt.Sprintf("Ты вакмен на %v%%", vacrate))
}
