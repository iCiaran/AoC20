package day_05

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartA(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "inputs/real_a.txt",
			want:  "965",
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
			want:  "524",
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

func TestDecode(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  int
	}{
		{
			input: "BFFFBBFRRR",
			want:  567,
		},
		{
			input: "FFFBBBFRRR",
			want:  119,
		},
		{
			input: "BBFFBBFRLL",
			want:  820,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := decode(test.input)

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
