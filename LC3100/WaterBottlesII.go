package LC3100

func MaxBottlesDrunk(numBottles int, numExchange int) int {

	out := numBottles
	emptyBottles :=  numBottles
	extraBottle := 0
	for emptyBottles>= numExchange {
		extraBottle++
		emptyBottles -= numExchange
		emptyBottles++
		numExchange++
	}

	return out + extraBottle
}