package app

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/nyan2d/menherabot/lol"
	tg "github.com/tucnak/telebot"
)

func (a *App) rotationCommand(m *tg.Message) {
	champs, err := lol.GetFreeChampPool()
	if err != nil {
		a.bot.Reply(m, "Не удалось получить список чампиков.")
		return
	}

	sort.Slice(champs, func(a, b int) bool {
		return strings.Compare(champs[a], champs[b]) < 0
	})
	a.bot.Reply(m, strings.Join(champs, "\r\n"))
}

func (a *App) rollCommand(m *tg.Message) {
	num := rand.Intn(99) + 1
	a.bot.Reply(m, fmt.Sprintf("%v", num))
}

func (a *App) vacmanCommand(m *tg.Message) {
	t := time.Now()
	vacrate := (t.Year() + int(t.Month()) + t.Day()*m.Sender.ID) % 100
	a.bot.Reply(m, fmt.Sprintf("Ты вакмен на %v%%", vacrate))
}
