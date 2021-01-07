package solution

func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	hold := -prices[0]
	sold := 0
	cd := 0
	for _, price := range prices {
		preSold := sold
		sold = hold + price
		hold = max(hold, cd-price)
		cd = max(cd, preSold)
	}
	return max(sold, cd)
}
