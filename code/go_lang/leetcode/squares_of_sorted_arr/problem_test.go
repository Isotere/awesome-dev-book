package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		Name string
		In   []int
		Out  []int
	}{
		{
			Name: "var 1",
			In:   []int{-4, -1, 0, 3, 10},
			Out:  []int{0, 1, 9, 16, 100},
		},
		{
			Name: "var 2",
			In:   []int{-7, -3, 2, 3, 11},
			Out:  []int{4, 9, 9, 49, 121},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			result := sortedSquares(test.In)

			assert.Equal(t, test.Out, result)
		})
	}
}
