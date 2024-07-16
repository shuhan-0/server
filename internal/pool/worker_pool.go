package pool

import (
	"sync"
)

type WorkerPool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

func NewWorkerPool(size int) *WorkerPool {
	pool := &WorkerPool{
		tasks: make(chan func(), size),
	}

	for i := 0; i < size; i++ {
		go pool.worker()
	}

	return pool
}

func (p *WorkerPool) worker() {
	for task := range p.tasks {
		task()
		p.wg.Done()
	}
}

func (p *WorkerPool) Submit(task func()) {
	p.wg.Add(1)
	p.tasks <- task
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
}

func (p *WorkerPool) Close() {
	close(p.tasks)
}
