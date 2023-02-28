package deleter

import (
	"context"
	"fmt"
	"github.com/cucumberjaye/url-shortener/internal/app/service"
	"github.com/cucumberjaye/url-shortener/models"
	"github.com/cucumberjaye/url-shortener/pkg/logger"
	"golang.org/x/sync/errgroup"
)

const workers = 5

type Deleter struct {
	repos service.DeleterRepository
	ch    chan []models.DeleteData
}

func New(repos service.DeleterRepository, ch chan []models.DeleteData) *Deleter {
	return &Deleter{
		repos: repos,
		ch:    ch,
	}
}

func (s *Deleter) Deleting() {
	fmt.Println("deleter running")

	g, _ := errgroup.WithContext(context.Background())

	shortCh := make(chan models.DeleteData)

	for i := 0; i < workers; i++ {
		g.Go(func() error {
			if err := s.repos.BatchDeleteURL(shortCh); err != nil {
				return err
			}

			return nil
		})
	}

	for data := range s.ch {
		for i := range data {
			shortCh <- data[i]
		}
	}

	if err := g.Wait(); err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
}
