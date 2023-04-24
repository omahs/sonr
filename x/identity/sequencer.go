package identity

import (
	"context"
	"log"
	"sync"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/crypto/mpc"
	"github.com/sonrhq/core/x/identity/controller"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/identity/types/models"
)

type Sequence interface {
	Next() *types.ClaimableWallet
	Add() error
}

type sequencer struct {
	sync.Mutex
	jobsQueue *Queue
	results   []*types.ClaimableWallet
}

func NewSequencer() Sequence {
	s := &sequencer{
		jobsQueue: NewQueue("NewWalletClaims"),
		results:   make([]*types.ClaimableWallet, 0),
	}
	wk := NewWorker(s.jobsQueue)
	go s.run(wk) // run the sequencer in the background
	return s
}

func (s *sequencer) Next() *types.ClaimableWallet {
	if len(s.results) == 0 {
		return nil
	}
	s.Lock()
	defer s.Unlock()
	w := s.results[0]
	s.results = s.results[1:]
	return w
}

func (s *sequencer) run(w *Worker) {
	for {
		done := w.DoWork()
		if done {
			log.Println("Worker done")
			break
		}
		continue
	}
}

func (s *sequencer) buildClaimableWallet() error {
	// Call Handler for keygen
	confs, err := mpc.Keygen(crypto.PartyID("current"))
	if err != nil {
		return err
	}

	var kss []models.KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			return err
		}
		ks, err := models.NewKeyshare(string(conf.ID), ksb, crypto.SONRCoinType, "primary")
		if err != nil {
			return err
		}
		kss = append(kss, ks)
	}
	cw, err := controller.NewWalletClaims("", kss)
	if err != nil {
		return err
	}
	s.results = append(s.results, cw.GetClaimableWallet())
	//	s.claimsCh <- cw
	return nil
}

func (s *sequencer) Add() error {
	// Call Handler for keygen
	confs, err := mpc.Keygen(crypto.PartyID("current"))
	if err != nil {
		return err
	}

	var kss []models.KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			return err
		}
		ks, err := models.NewKeyshare(string(conf.ID), ksb, crypto.SONRCoinType, "primary")
		if err != nil {
			return err
		}
		kss = append(kss, ks)
	}
	cw, err := controller.NewWalletClaims("", kss)
	if err != nil {
		return err
	}
	action := func() error {
		s.results = append(s.results, cw.GetClaimableWallet())
		return nil
	}
	job := Job{
		Name:   "Build Claimable Wallet",
		Action: action,
	}
	s.jobsQueue.AddJob(job)
	return nil
}

// Queue holds name, list of jobs and context with cancel.
type Queue struct {
	name   string
	jobs   chan Job
	ctx    context.Context
	cancel context.CancelFunc
}

// Job - holds logic to perform some operations during queue execution.
type Job struct {
	Name   string
	Action func() error // A function that should be executed when the job is running.
}

// NewQueue instantiates new queue.
func NewQueue(name string) *Queue {
	ctx, cancel := context.WithCancel(context.Background())

	return &Queue{
		jobs:   make(chan Job),
		name:   name,
		ctx:    ctx,
		cancel: cancel,
	}
}

// AddJobs adds jobs to the queue and cancels channel.
func (q *Queue) AddJobs(jobs []Job) {
	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, job := range jobs {
		// Goroutine which adds job to the queue.
		go func(job Job) {
			q.AddJob(job)
			wg.Done()
		}(job)
	}

	go func() {
		wg.Wait()
		// Cancel queue channel, when all goroutines were done.
		q.cancel()
	}()
}

// AddJob sends job to the channel.
func (q *Queue) AddJob(job Job) {
	q.jobs <- job
	log.Printf("New job %s added to %s queue", job.Name, q.name)
}

// Run performs job execution.
func (j Job) Run() error {
	log.Printf("Job running: %s", j.Name)

	err := j.Action()
	if err != nil {
		return err
	}

	return nil
}

// Worker responsible for queue serving.
type Worker struct {
	Queue *Queue
}

// NewWorker initializes a new Worker.
func NewWorker(queue *Queue) *Worker {
	return &Worker{
		Queue: queue,
	}
}

// DoWork processes jobs from the queue (jobs channel).
func (w *Worker) DoWork() bool {
	for {
		select {
		// if context was canceled.
		case <-w.Queue.ctx.Done():
			log.Printf("Work done in queue %s: %s!", w.Queue.name, w.Queue.ctx.Err())
			return true
		// if job received.
		case job := <-w.Queue.jobs:
			err := job.Run()
			if err != nil {
				log.Print(err)
				continue
			}
		}
	}
}
