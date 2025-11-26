package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateMonthsToDeplete(t *testing.T) {
	tests := []struct {
		amountSaved       float64
		monthlyWithdrawal float64
		monthlyReturn     float64
		expectedMonths    float64
		expectError       bool
	}{
		{150000, 10000, 1, 15.338527921527762, false},
		{50000, 10000, 1, 4.162656342548076, false},
		{100000, 5000, 0.5, 20.128823382622716, false},
		{200000, 15000, 2, 14.666760025970527, false},
		{1000000, 2000, 1, 0, true}, // monthly withdrawal too low/high combination
		{50000, 60000, 1, 0, true},  // withdrawal larger than balance
		{100000, 5000, -1, 0, true}, // negative monthly return
	}

	for _, test := range tests {
		months, err := CalculateMonthsToDeplete(test.amountSaved, test.monthlyWithdrawal, test.monthlyReturn)
		require.Equal(t, test.expectError, err != nil)
		require.Equal(t, test.expectedMonths, months)
	}
}
