package main

import (
	"fmt"
	"testing"
)

// - struct test, fields: data []int, answer int
// - tests := []test{[]int{}, int}
// - range tests

type test struct {
	data   []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		test{data: []int{1, 2, 3}, answer: 6},
		test{[]int{10, 11, 12}, 33},
		test{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		teste := Soma(v.data...)
		resultado := v.answer
		if teste != v.answer {
			t.Error("Expected:", resultado, "Got:", teste)
		}
	}
}

func ExampleSoma() {
	fmt.Println(Soma(4, 2, 1))
	fmt.Println(Soma(4, 2, 2))
	fmt.Println(Soma(4, 2, 3))
	// Output:
	// 7
	// 8
	// 9
}

func TestSoma(t *testing.T) {
	teste := Soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func TestSoma2(t *testing.T) {
	teste := Soma2(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func TestMultiplica(t *testing.T) {
	teste := Multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func TestMultiplica2(t *testing.T) {
	teste := Multiplica2(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(1, 2,3,4,5,6,7,8,9,10)
		Soma(10,11,12,13,14,15,16,17,18,19,20)
	}
}

func BenchmarkSoma2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma2(1, 2,3,4,5,6,7,8,9,10)
		Soma2(10,11,12,13,14,15,16,17,18,19,20)
	}
}

func BenchmarkMultiplica(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplica(2,3,4,5,6,7,8,9,10)
		Multiplica(10,11,12,13,14,15,16,17,18,19,20)
	}
}

func BenchmarkMultiplica2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplica2(2,3,4,5,6,7,8,9,10)
		Multiplica2(10,11,12,13,14,15,16,17,18,19,20)
	}
}