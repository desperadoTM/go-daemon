package daemon

import (
	"context"
	"log"
	"os"
	"sync"
	"time"
)

type Server interface {
	Start(ctx context.Context, d time.Duration, msg string, wg sync.WaitGroup)
	Stop()
}

type SampleDaemon struct {
}

func New() SampleDaemon {
	return SampleDaemon{}
}

func (s *SampleDaemon) Start(ctx context.Context, d time.Duration, msg string, wg sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-time.After(d):
			log.Println(msg)
			//логика обработки
			baseProcess()
		case <-ctx.Done():
			log.Print(ctx.Err())
			return
		}
	}
}

func (s *SampleDaemon) Stop() {
	log.Println("stopping")
	os.Exit(1)
}

// Пока пусть будет спать
func baseProcess() {
	log.Println("sleep step 1...")
	time.Sleep(5 * time.Second)
	log.Println("sleep step 2...")
}
