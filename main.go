package main

import "fmt"

func main() {
	var amountSaved, monthlyWithdrawal, monthlyReturnPercent float64

	fmt.Println("Enter the amount saved:")
	fmt.Scan(&amountSaved)

	fmt.Println("Enter the fixed monthly withdrawal:")
	fmt.Scan(&monthlyWithdrawal)

	fmt.Println("Enter the monthly return (as a decimal, e.g. 1 for 1%):")
	fmt.Scan(&monthlyReturnPercent)

	months, err := CalculateMonthsToDeplete(amountSaved, monthlyWithdrawal, monthlyReturnPercent)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("It will take %.2f months to deplete the savings.\n", months)
}

// CalculateMonthsToDeplete returns the number of months until the savings
// is depleted given an initial amount, a fixed monthly withdrawal and a
// monthly return expressed as a percent (e.g. pass 1 for 1%). The function
// preserves the original calculation logic; it returns an error for invalid
// parameter combinations.
func CalculateMonthsToDeplete(amountSaved, monthlyWithdrawal, monthlyReturnPercent float64) (float64, error) {
	monthlyReturn := monthlyReturnPercent / 100
	if monthlyReturn <= 0 {
		return 0, fmt.Errorf("monthly return must be greater than 0")
	}

	if monthlyWithdrawal >= (amountSaved + (amountSaved * monthlyReturn)) {
		return 0, fmt.Errorf("monthly withdrawal is too large to deplete savings")
	}

	if amountSaved*monthlyReturn >= monthlyWithdrawal {
		return 0, fmt.Errorf("savings will never deplete with the given parameters")
	}

	months := 0.0
	for amountSaved > 0 {
		amountSaved += amountSaved * monthlyReturn
		amountSaved -= monthlyWithdrawal
		months++
	}
	// adjust for the final partial month
	months--
	months += amountSaved / (monthlyWithdrawal - (amountSaved * monthlyReturn))

	return months, nil
}
