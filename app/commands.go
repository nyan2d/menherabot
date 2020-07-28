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
	"github.com/nyan2d/menherabot/util"
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
	ti := util.TimeToInt(time.Now().UTC().Add(time.Hour * 3))
	ti += m.Sender.ID
	rate := ((1664525*ti + 1013904223) % 2147483647) % 101
	a.bot.Reply(m, fmt.Sprintf("Ты вакмен на %v%%", rate))
}
