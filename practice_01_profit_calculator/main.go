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
	fmt.Println("Earning before tax: ", earningBeforeTax)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)

}