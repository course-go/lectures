package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type (
	data struct {
		id int
	}
	result struct{}
)

type dataResult struct {
	result result
	err    error
}

func main() {
	data := []data{
		{
			id: 1,
		},
		{
			id: 2,
		},
		{
			id: 3,
		},
	}

	_ = processData(context.Background(), data)
}

// START OMIT

func processData(ctx context.Context, data []data) (err error) {
	var wg sync.WaitGroup
	resultCh := make(chan dataResult, len(data))
	for _, d := range data {
		wg.Go(func() {
			runWorker(ctx, d, resultCh)
		})
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	for result := range resultCh {
		if result.err != nil {
			err = errors.Join(err, result.err)
			continue
		}
		// Process result data...
	}
	if err != nil {
		return fmt.Errorf("some of the workers failed to process data: %w", err)
	}
	return nil
}

// END OMIT

func runWorker(ctx context.Context, d data, resultCh chan<- dataResult) {
	t := time.NewTimer(time.Duration(rand.Int()%500) * time.Millisecond)

	select {
	case <-t.C:
		fmt.Println("Worker", d.id, "executed workload!")
	case <-ctx.Done():
		fmt.Println("Worker", d.id, "cancelled!")
	}

	resultCh <- dataResult{}
}
