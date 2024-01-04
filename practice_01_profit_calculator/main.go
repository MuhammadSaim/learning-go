package main

import "fmt"

func main(){
	var revenue float64
	var expenses float64
	var taxRate float64

	// get the input for revenue
	fmt.Print("Revenue: ");
	fmt.Scan(&revenue)

	// get the input for expenses
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	// get the input for taxRate
	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

	// find the earnings before the taxes
	earningBeforeTax := revenue - expenses

	// find the actual profit after taxes
	profit := earningBeforeTax * (1 - taxRate/100)

	// find the ratio
	ratio := earningBeforeTax / profit


	// print the earning before tax
	fmt.Printf("Earning before tax: %.2f\n", earningBeforeTax)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

}