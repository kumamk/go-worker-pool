package workerpool

import "log"

// worker pool behaviour
type Workerpool interface {
	Init()
	AddJob(job func())
}

type workerpool struct {
	maxWorker int
	jobQueue  chan func()
}

func NewWorkerPool(maxWorker int) Workerpool {
	wp := &workerpool{
		maxWorker: maxWorker,
		jobQueue:  make(chan func()),
	}
	return wp
}

func (wp *workerpool) Init() {
	for i := 0; i < wp.maxWorker; i++ {
		workerId := i + 1
		log.Printf("Init worker %d", workerId)

		go func(id int) {
			for job := range wp.jobQueue {
				log.Printf("Worker %d started the job", id)
				job()
				log.Printf("Worker %d finished the job", id)
			}

		}(workerId)
	}
}

func (wp *workerpool) AddJob(job func()) {
	wp.jobQueue <- job
}
