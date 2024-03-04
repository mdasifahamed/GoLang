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

// multiple fucntion return
func actualInterestAfterTaxdDeduction(amount int, rate float32, tax float32) (monthlyProfitAftreTax float32, yearlyProfitAfterTax float32, principalAndProfitAfterTaxYearly float32) {
	yearlyProfitAfterTax = calculateInterestYearly(amount, rate) - ((calculateInterestYearly(amount, rate) * tax) / 100.0)
	monthlyProfitAftreTax = calculateInterestMonthly(amount, rate) - ((calculateInterestMonthly(amount, rate) * tax) / 100.0)
	principalAndProfitAfterTaxYearly = yearlyProfitAfterTax + float32(amount)
	return
}

// Variadic functions takes dynamic number of parameter of a single data type

// craete int Array

func createArrayofevenNumbers(nums ...int) []int {
	var storeEven []int
	// using range to determine length
	for _, el := range nums { // "_" this omits the index we don't nedd index here
		if el%2 == 0 {
			storeEven = append(storeEven, el)
		}
	}
	return storeEven
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

	// amount is 1524700 rate 10% tax 5%

	var monthly, yearly, principleAndProfit = actualInterestAfterTaxdDeduction(amountToinvest, rate, 5)

	fmt.Println("monthly interest after tax", monthly)
	fmt.Println("yearly interest after tax", yearly)
	fmt.Println("Principla and Profit Yearly  after tax", principleAndProfit)

	// using fucntion create a array of even number from differents numbers
	evens := createArrayofevenNumbers(1, 5, 6, 79, 12, 64, 89, 74, 58, 69, 32, 14, 89, 78, 98, 65, 23, 14, 10, 200, 64, 78, 39, 96)
	fmt.Println("Set Of Even Numbers", evens)

}
