package synctest_test

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
)

// AFTERFUNC START OMIT

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

// AFTERFUNC END OMIT
// TIME START OMIT

func TestContextWithTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(t.Context(), timeout)
		defer cancel()

		time.Sleep(timeout - time.Nanosecond)
		synctest.Wait()
		if err := ctx.Err(); err != nil {
			t.Fatalf("before timeout: ctx.Err() = %v, want nil\n", err)
		}

		time.Sleep(time.Nanosecond)
		synctest.Wait()
		if err := ctx.Err(); err != context.DeadlineExceeded {
			t.Fatalf("after timeout: ctx.Err() = %v, want DeadlineExceeded\n", err)
		}
	})
}

// TIME END OMIT
