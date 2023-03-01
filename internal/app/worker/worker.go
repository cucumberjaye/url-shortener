package worker

import (
	"context"
	"fmt"
	"github.com/cucumberjaye/url-shortener/models"
	"sync"
)

const workers = 5

type DeleterRepository interface {
	BatchDeleteURL(short, id string) error
}

type Worker struct {
	wg         *sync.WaitGroup
	cancelFunc context.CancelFunc
	repo       DeleterRepository
	ch         chan models.DeleteData
}

func New(repo DeleterRepository, ch chan models.DeleteData) *Worker {
	return &Worker{
		wg:   new(sync.WaitGroup),
		repo: repo,
		ch:   ch,
	}
}

func (w *Worker) Start(pctx context.Context) {
	fmt.Println("worker running")
	ctx, cancelFunc := context.WithCancel(pctx)
	w.cancelFunc = cancelFunc
	for i := 0; i < workers; i++ {
		w.wg.Add(1)
		go w.spawnWorkers(ctx)
	}
}

func (w *Worker) Stop() {
	w.cancelFunc()
	w.wg.Wait()
}

func (w *Worker) spawnWorkers(ctx context.Context) {
	defer w.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-w.ch:
			if err := w.repo.BatchDeleteURL(data.ShortURL, data.ID); err != nil {
				fmt.Println(err)
			}
		}
	}
}
