package day_13

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
			input: "inputs/test_01.txt",
			want:  "295",
		},
		{
			input: "inputs/real_a.txt",
			want:  "2845",
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
			input: "inputs/test_01.txt",
			want:  "1068781",
		},
		{
			input: "inputs/test_02.txt",
			want:  "3417",
		},
		{
			input: "inputs/test_03.txt",
			want:  "754018",
		},
		{
			input: "inputs/test_04.txt",
			want:  "779210",
		},
		{
			input: "inputs/test_05.txt",
			want:  "1261476",
		},
		{
			input: "inputs/test_06.txt",
			want:  "1202161486",
		},
		{
			input: "inputs/real_b.txt",
			want:  "487905974205117",
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
