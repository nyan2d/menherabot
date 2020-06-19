package app

import (
	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"time"

	"github.com/nyan2d/menherabot/config"
	tg "github.com/tucnak/telebot"
)

type App struct {
	bot *tg.Bot
	leagueClient *golio.Client
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
		leagueClient: golio.NewClient(cfg.LeagueToken, golio.WithRegion(api.RegionEuropeWest)),
	}

	a.bindHandlers()

	return a
}

func (a *App) bindHandlers() {
	a.bot.Handle("/roll", a.rollCommand)
	a.bot.Handle("/rotate", a.champikiCommand)
	a.bot.Handle("/vacman", a.vacmanCommand)
}

func (a *App) Run() {
	a.bot.Start()
}
