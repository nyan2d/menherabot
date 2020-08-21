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
		Token:  cfg.BotToken,
		Poller: NewLimiter(3 * time.Second).Poller(),
	}

	if b, err := tg.NewBot(settings); err != nil {
		panic(err)
	} else {
		a := &App{
			bot: b,
		}
		a.bindHandlers()
		return a
	}
}

func (a *App) bindHandlers() {
	a.bot.Handle("/roll", a.rollCommand)
	a.bot.Handle("/pick", a.pickCommand)
	a.bot.Handle("/rotation", a.rotationCommand)
	a.bot.Handle("/fart", a.fartCommand)
}

func (a *App) Run() {
	a.bot.Start()
}
