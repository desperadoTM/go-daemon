package go_deamon

import (
	"context"
	"go-deamon/daemon"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup
	testMessage := "hello from daemon"

	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case <-c:
			log.Println("Получено прерывание, ждем завершение процесса перед выходом")
			cancel()
		case <-ctx.Done():

		}
	}()
	svr := daemon.New()
	svr.Start(ctx, 5*time.Second, testMessage, wg)
	wg.Wait()
	svr.Stop()
}
