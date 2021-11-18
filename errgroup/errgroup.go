package errgroup

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func Run() {
	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		time.Sleep(1 * time.Second)
		log.Println("sleep 1 second")
		return errors.New("sleep 1 second")
	})

	eg.Go(func() error {
		time.Sleep(2 * time.Second)
		log.Println("sleep 2 second")
		return errors.New("sleep 2 second")
	})

	eg.Go(func() error {
		time.Sleep(5 * time.Second)
		log.Println("sleep 5 second")
		return errors.New("sleep 5 second")
	})

	eg.Go(func() error {
		<-ctx.Done()
		log.Println("ctx Done")
		return errors.New("ctx Done")
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		fmt.Println(err)
	}
}
