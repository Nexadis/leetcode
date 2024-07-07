package main

func numWaterBottles(numBottles int, numExchange int) int {
	drinked := 0
	withWater := numBottles
	for withWater != 0 {
		drinked += withWater
		withWater = numBottles / numExchange
		numBottles = withWater + numBottles%numExchange
	}
	return drinked
}
