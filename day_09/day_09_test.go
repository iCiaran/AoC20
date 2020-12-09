package day_09

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartA(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		preamble int
		input    string
		want     string
	}{
		{
			preamble: 5,
			input:    "inputs/test_01.txt",
			want:     "127",
		},
		{
			preamble: 25,
			input:    "inputs/real_a.txt",
			want:     "105950735",
		},
	}

	day := New()
	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			day.preamble = test.preamble
			got := day.PartA(test.input)

			assert.Equal(test.want, got)
		})
	}
}

func TestPartB(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		preamble int
		input    string
		want     string
	}{
		{
			preamble: 5,
			input:    "inputs/test_01.txt",
			want:     "62",
		},
		{
			preamble: 25,
			input:    "inputs/real_b.txt",
			want:     "13826915",
		},
	}
	day := New()
	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			day.preamble = test.preamble
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
