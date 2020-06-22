package app

import (
	"time"

	tg "github.com/tucnak/telebot"
)

type Limiter struct {
	users  map[int]time.Time
	poller *tg.MiddlewarePoller
}

func NewLimiter(rate time.Duration) *Limiter {
	longpoller := &tg.LongPoller{
		Timeout: 10 * time.Second,
	}

	r := Limiter{
		users: map[int]time.Time{},
	}

	r.poller = tg.NewMiddlewarePoller(longpoller, func(u *tg.Update) bool {
		if u.Message == nil {
			return true
		}
		ti, ok := r.users[u.Message.Sender.ID]
		if !ok || ti.Before(time.Now()) {
			r.users[u.Message.Sender.ID] = time.Now().Add(rate)
			return true
		}
		return false
	})

	return &r
}

func (l *Limiter) Poller() *tg.MiddlewarePoller {
	return l.poller
}
