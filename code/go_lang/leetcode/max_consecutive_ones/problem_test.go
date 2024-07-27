package max_consecutive_ones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		Name string
		In   []int
		Out  int
	}{
		{
			Name: "empty",
			In:   []int{},
			Out:  0,
		},
		{
			Name: "all zeros",
			In:   []int{0, 0},
			Out:  0,
		},
		{
			Name: "all ones",
			In:   []int{1, 1},
			Out:  2,
		},
		{
			Name: "variant 1",
			In:   []int{1, 1, 0, 1, 1, 1},
			Out:  3,
		},
		{
			Name: "variant 2",
			In:   []int{1, 0, 1, 1, 0, 1},
			Out:  2,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			result := findMaxConsecutiveOnes(test.In)

			assert.Equal(t, test.Out, result)
		})
	}
}
