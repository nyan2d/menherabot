package app

import (
	"fmt"
	"math/rand"

	tg "github.com/tucnak/telebot"
)

func (a *App) rollCommand(m *tg.Message) {
	num := rand.Intn(99) + 1
	a.bot.Reply(m, fmt.Sprintf("%v", num))
}

func (a *App) vacmanCommand(m *tg.Message) {
	vacrate := m.Sender.ID % 100
	a.bot.Reply(m, fmt.Sprintf("Ты вакмен на %v%%", vacrate))
}
