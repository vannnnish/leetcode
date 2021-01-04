package 买卖股票最佳时期

func maxProfit(prices []int) int {
	profit := 0
	length := len(prices)
	haveGoods := false
	goodsPrice := 0
	if length >= 2 {
		for i := 0; i < length; i++ {
			if i <= length-2 {
				if !haveGoods && prices[i] < prices[i+1] {
					haveGoods = true
					goodsPrice = prices[i]
				}
			}

			if haveGoods {
				// 开始跌价， 则出售
				if i+1 == length || prices[i+1] < prices[i] {
					profit = profit + prices[i] - goodsPrice
					haveGoods = false
				}
			}
		}
	}
	return profit
}
