package app

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
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
	limit := 99
	limstr := regexp.MustCompile(`^[1-9]\d{0,4}$`).FindString(m.Payload)
	if len(limstr) > 0 {
		limit, _ = strconv.Atoi(limstr)
	}
	num := rand.Intn(limit) + 1
	a.bot.Reply(m, fmt.Sprintf("%v", num))
}

func (a *App) pickCommand(m *tg.Message) {
	if len(m.Payload) == 0 {
		a.bot.Reply(m, "Ты дебил?")
		return
	}
	list := strings.Split(m.Payload, " ")
	if len(list) < 2 {
		a.bot.Reply(m, "Пошел нах!")
		return
	}
	choise := list[rand.Intn(len(list))]
	a.bot.Reply(m, strings.Replace(choise, "_", " ", -1))
}

func (a *App) vacCommand(m *tg.Message) {
	t := time.Now()
	vacrate := (t.Year() + int(t.Month()) + t.Day()*m.Sender.ID) % 100
	a.bot.Reply(m, fmt.Sprintf("Ты вакмен на %v%%", vacrate))
}
