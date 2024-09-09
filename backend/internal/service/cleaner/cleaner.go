package cleaner

import (
	"context"
	"errors"
	"time"

	gocron "github.com/go-co-op/gocron/v2"

	tokenrepo "travelgo/internal/repository/token"
)

type Cleaner struct {
	repo *tokenrepo.Repository

	scheduler gocron.Scheduler
}

func New(repo *tokenrepo.Repository) *Cleaner {
	cl := &Cleaner{repo: repo}

	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	cl.scheduler = s

	return cl
}

func (c *Cleaner) Start(ctx context.Context, interval time.Duration) error {
	_, err := c.scheduler.NewJob(
		gocron.DurationJob(
			interval,
		),
		gocron.NewTask(
			c.repo.CleanExpiredTokens,
			ctx,
		),
	)
	if err != nil {
		return errors.New("NewJob: " + err.Error())
	}

	c.scheduler.Start()
	return nil
}

func (c *Cleaner) Stop() error {
	return c.scheduler.Shutdown()
}
