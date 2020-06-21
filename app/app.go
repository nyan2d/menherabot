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

	a.bindHandlers()

	return a
}

func (a *App) bindHandlers() {
	a.bot.Handle("/roll", a.rollCommand)
	a.bot.Handle("/pick", a.pickCommand)
	a.bot.Handle("/rotation", a.rotationCommand)
	a.bot.Handle("/vac", a.vacCommand)
}

func (a *App) Run() {
	a.bot.Start()
}
