package day_17

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartA(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "inputs/real_a.txt",
			want:  "A",
		},
	}

    day := New()
	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := day.PartA(test.input)

			assert.Equal(test.want, got)
		})
	}
}

func TestPartB(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "inputs/real_b.txt",
			want:  "B",
		},
	}
    day := New()
	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := day.PartB(test.input)

			assert.Equal(test.want, got)
		})
	}
}

func BenchmarkPartA(b *testing.B) {
    day := New()
	for n := 0; n < b.N; n++ {
		day.PartA("inputs/real_a.txt")
	}
}

func BenchmarkPartB(b *testing.B) {
    day := New()
	for n := 0; n < b.N; n++ {
		day.PartB("inputs/real_b.txt")
	}
}
