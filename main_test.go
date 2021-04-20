package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test1(t *testing.T) {
	matrix := [][]int{
		{10, 20},
		{30, 200},
		{400, 50},
		{30, 20},
	}
	expected := 110
	actual := MinCost(matrix)
	require.Equal(t, expected, actual)
}
