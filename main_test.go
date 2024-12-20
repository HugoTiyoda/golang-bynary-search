package main

import (
	"math/rand"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := generateArray(100)
	target := 42

	linearResult := linearSearch(arr, target)
	if linearResult != target {
		t.Errorf("LinearSearch falhou: esperado %d, obtido %d", target, linearResult)
	}

	binaryResult := binarySearch(arr, target)
	if binaryResult != target {
		t.Errorf("BinarySearch falhou: esperado %d, obtido %d", target, binaryResult)
	}
}

// Benchmarks
func BenchmarkLinearSearch(b *testing.B) {
	arr := generateArray(100)
	target := 99 // Pior caso

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linearSearch(arr, target)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	arr := generateArray(100)
	target := 99

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binarySearch(arr, target)
	}
}

func BenchmarkLinearSearchRandom(b *testing.B) {
	arr := generateArray(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		target := arr[rand.Intn(len(arr))]
		linearSearch(arr, target)
	}
}

func BenchmarkBinarySearchRandom(b *testing.B) {
	arr := generateArray(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		target := arr[rand.Intn(len(arr))]
		binarySearch(arr, target)
	}
}
