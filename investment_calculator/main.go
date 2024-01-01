package main

import (
	"fmt"
	"math"
)

func main(){

	// total amount user is investing
	var investmentAmount float64 = 1000

	// return rate of an amount in percentage
	var expectedReturnRate = 5.5

	// number of years to to hold the investement
	var years float64 = 10

	// investment formula
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	// print out the future value
	fmt.Print(futureValue)
}