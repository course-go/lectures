package main

import "testing"

// START OMIT

func BenchmarkReset(b *testing.B) {
	// init benchmark
	b.ResetTimer()
	for range b.N {
		// benchmark
	}
}

func BenchmarkLoop(b *testing.B) {
	for b.Loop() { // resets timer
		// benchmark
	}
}

func BenchmarkStopStart(b *testing.B) {
	for range b.N {
		b.StopTimer()
		// init iteration benchmark
		b.StartTimer()
		// benchmark
	}
}

// END OMIT
