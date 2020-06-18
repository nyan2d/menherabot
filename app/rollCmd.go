package app

import (
	"fmt"
	"math/rand"

	tg "github.com/tucnak/telebot"
)

func (a *App) rollCmd(m *tg.Message) {
	num := rand.Intn(99) + 1
	a.bot.Reply(m, fmt.Sprintf("%v", num))
}
