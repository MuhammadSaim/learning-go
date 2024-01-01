package main

import (
	"fmt"
	"math"
)

func main(){

	// set the infelation rate per year
	const infelationRate = 2.5

	// An investment value user want to invest
	var investmentAmount float64;

	// numbers of years to hold investment
	var years float64

	// return rate expected in percentage
	expectedReturnRate := 5.5

	// get the investment value from user
	fmt.Print("Investment value: ")
	fmt.Scan(&investmentAmount)

	// get the value of an expectedReturnRate
	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// get the years value from the user
	fmt.Print("Enter total years: ")
	fmt.Scan(&years)

	// investment formula to get future value
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	// formula to get future real value
	var futureRealValue = futureValue * math.Pow(1+infelationRate/100, years)

	// outputing the values
	fmt.Println("Future Value: ", futureValue)
	fmt.Println("Future Real Value: ", futureRealValue)

}