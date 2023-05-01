package worker

import (
	"context"
	"fmt"
	"sync"

	"github.com/cucumberjaye/url-shortener/models"
)

// количество горутин, обрабатывающий запросы на удаление
const workers = 5

// интерфейс хранилища для удаления запросов
type DeleterRepository interface {
	BatchDeleteURL(short, id string) error
}

// структура, отвечающая за удаление ссылок
type Worker struct {
	wg         *sync.WaitGroup
	cancelFunc context.CancelFunc
	repo       DeleterRepository
	ch         chan models.DeleteData
}

// создаем Worker
func New(repo DeleterRepository, ch chan models.DeleteData) *Worker {
	return &Worker{
		wg:   new(sync.WaitGroup),
		repo: repo,
		ch:   ch,
	}
}

// запускем воркер
func (w *Worker) Start(pctx context.Context) {
	fmt.Println("worker running")
	ctx, cancelFunc := context.WithCancel(pctx)
	w.cancelFunc = cancelFunc
	for i := 0; i < workers; i++ {
		w.wg.Add(1)
		go w.spawnWorkers(ctx)
	}
}

// останавливаем воркер
func (w *Worker) Stop() {
	w.cancelFunc()
	w.wg.Wait()
}

// запускаем горутины, которые ждут данные из канала для удаления из хранилища
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
