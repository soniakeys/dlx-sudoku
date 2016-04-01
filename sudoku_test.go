package main

import "testing"

const (
	hard = "8..........36......7..9.2...5...7.......457.....1...3...1....68..85...1..9....4.."
	pb   = ".......12........3..23..4....18....5.6..7.8.......9.....85.....9...4.5..47...6..."
)

func Test(t *testing.T) {
	if Solve(hard) == "" {
		t.Fatal("nope")
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Solve(pb)
	}
}
