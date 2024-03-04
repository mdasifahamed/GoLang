package main

import "fmt"

// function without return value and parameter

func displaygo() {
	fmt.Println("hello go")
}

// parametrized function
func calculateInterestYearly(amount int, rate float32) float32 {
	totalInterest := (float32(amount) * rate) / 100.0
	return totalInterest
}

func calculateInterestMonthly(amount int, rate float32) float32 {
	monthlyProfit := calculateInterestYearly(amount, rate) / 12.0
	return monthlyProfit
}
func main() {
	displaygo()

	// 1524700 for 10% Yearly Interrest Would be
	amountToinvest := 1524700
	var rate float32 = 10.0

	profit := calculateInterestYearly(amountToinvest, rate)
	fmt.Println("Profit generated in a Year", profit)
	monthlyReturn := calculateInterestMonthly(amountToinvest, rate)
	fmt.Println("Profit generated in a month", monthlyReturn)

}
