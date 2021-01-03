package day_18

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
			want:  "51",
		},
		{
			input: "inputs/test_02.txt",
			want:  "26",
		},
		{
			input: "inputs/test_03.txt",
			want:  "437",
		},
		{
			input: "inputs/test_04.txt",
			want:  "12240",
		},
		{
			input: "inputs/test_05.txt",
			want:  "13632",
		},
		{
			input: "inputs/real_a.txt",
			want:  "67800526776934",
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
			want:  "51",
		},
		{
			input: "inputs/test_02.txt",
			want:  "46",
		},
		{
			input: "inputs/test_03.txt",
			want:  "1445",
		},
		{
			input: "inputs/test_04.txt",
			want:  "669060",
		},
		{
			input: "inputs/test_05.txt",
			want:  "23340",
		},
		{
			input: "inputs/real_b.txt",
			want:  "340789638435483",
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
