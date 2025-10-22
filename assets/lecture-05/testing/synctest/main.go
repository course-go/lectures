package main

import (
	"context"
	"testing"
	"testing/synctest"
)

// START OMIT

func TestContextAfterFunc(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(t.Context())
		var afterFuncCalled bool
		context.AfterFunc(ctx, func() {
			afterFuncCalled = true
		})

		synctest.Wait()
		if afterFuncCalled {
			t.Fatalf("before context is canceled: AfterFunc called")
		}

		cancel()
		synctest.Wait()
		if !afterFuncCalled {
			t.Fatalf("before context is canceled: AfterFunc not called")
		}
	})
}

// END OMIT
