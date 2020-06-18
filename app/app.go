package app

import (
	"time"

	"github.com/nyan2d/menherabot/config"
	tg "github.com/tucnak/telebot"
)

type App struct {
	bot *tg.Bot
}

func NewApp(cfg *config.Config) *App {
	settings := tg.Settings{
		Token: cfg.BotToken,
		Poller: &tg.LongPoller{
			Timeout: 10 * time.Second,
		},
	}

	b, err := tg.NewBot(settings)
	if err != nil {
		panic(err)
	}

	a := &App{
		bot: b,
	}

	a.bindHandlers(b)

	return a
}

func (a *App) bindHandlers(b *tg.Bot) {
	a.bot.Handle("/roll", a.rollCmd)
}

func (a *App) Run() {
	a.bot.Start()
}
