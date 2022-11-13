package app

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"
)

type (
	Runner interface {
		Start() error
		Stop(ctx context.Context) error
	}

	Application struct {
		runners []Runner
	}
)

func New(runners ...Runner) *Application {
	return &Application{
		runners: runners,
	}
}

func (a Application) Start(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	for _, s := range a.runners {
		go func(runner Runner) {
			if err := runner.Start(); err != nil {
				log.Fatalln(err)
			}
		}(s)
	}

	fmt.Println("all systems go!")

	<-ctx.Done()
	stop()

	fmt.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, s := range a.runners {
		if err := s.Stop(ctx); err != nil {
			log.Println("oops...", err)
		}
	}

	fmt.Println("good bye")
}
