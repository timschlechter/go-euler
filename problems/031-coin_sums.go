/*
In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/
package problems

func howManyWays(amount int, coins []int) int {
	if len(coins) == 0 {
		return 0
	}
	coin := coins[0]
	if coin > amount {
		return howManyWays(amount, coins[1:])
	}
	if coin < amount {
		return howManyWays(amount-coin, coins) + howManyWays(amount, coins[1:])
	}
	return 1 + howManyWays(amount, coins[1:])
}

func Problem031() int {
	return howManyWays(200, []int{200, 100, 50, 20, 10, 5, 2, 1})
}
