package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	tf, err := openTraceFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed creating assets: %v", err)
		os.Exit(1)
	}

	defer func() {
		_ = tf.Close()
	}()

	err = runApp(context.Background(), tf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed running app: %v", err)
		os.Exit(1)
	}
}

// START OMIT

func runApp(ctx context.Context, traceFile io.WriteCloser) error {
	err := trace.Start(traceFile)
	if err != nil {
		return fmt.Errorf("failed starting tracing: %w", err)
	}
	defer trace.Stop()

	trace.WithRegion(ctx, "makeCappuccino", func() {
		// orderID allows to identify a specific order
		// among many cappuccino order region records.
		trace.Log(ctx, "orderID", orderID())

		trace.WithRegion(ctx, "steamMilk", steamMilk)
		trace.WithRegion(ctx, "extractCoffee", extractCoffee)
		trace.WithRegion(ctx, "mixMilkCoffee", mixMilkCoffee)
	})

	return nil
}

// END OMIT

func openTraceFile() (wc io.WriteCloser, err error) {
	file, err := os.Create("trace.out")
	if err != nil {
		return nil, fmt.Errorf("failed opening trace file: %w", err)
	}

	return file, nil
}

func orderID() string {
	return "0f70e7ff-2f01-4006-bffa-1091716b43d3"
}

func steamMilk() {
	time.Sleep(2 * time.Microsecond)
}

func extractCoffee() {
	time.Sleep(3 * time.Microsecond)
}

func mixMilkCoffee() {
	time.Sleep(1 * time.Microsecond)
}
