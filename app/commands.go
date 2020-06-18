package app

import (
	"fmt"
	"math/rand"

	tg "github.com/tucnak/telebot"
)

func (a *App) rollCommand(m *tg.Message) {
	a.bot.Reply(m, fmt.Sprintf("%v", rand.Intn(99) + 1))
}

func (a *App) vakmanCommand(m *tg.Message) {
	x := m.Sender.ID % 100
	_, err := a.bot.Reply(m, fmt.Sprintf("ты вакмен на %v%%", x))
	if err != nil {
		fmt.Printf("Пиздячка: %v", m.Sender.ID)
	}
}
