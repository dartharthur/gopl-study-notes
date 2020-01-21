package main

import (
	"testing"
)

func benchmarkEcho(b *testing.B, size int) {
	inputData := buildInputData(size, []string{"hello", "darkness", "my", "old", "friend"})
	for i := 0; i < b.N; i++ {
		echo(inputData)
	}
}

func BenchmarkEcho1(b *testing.B)    { benchmarkEcho(b, 1) }
func BenchmarkEcho10(b *testing.B)   { benchmarkEcho(b, 10) }
func BenchmarkEcho100(b *testing.B)  { benchmarkEcho(b, 100) }
func BenchmarkEcho1000(b *testing.B) { benchmarkEcho(b, 1000) }
