/*
LeetCode Problem #901: Online Stock Span
Difficulty: Medium

Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.
*/

package LC901

type pair struct {
	price int
	span  int
}
type StockSpanner struct {
	stocks []pair
}

func Constructor() StockSpanner {
	return StockSpanner{stocks: make([]pair, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	locSpan := 1
	for len(this.stocks) > 0 && this.stocks[len(this.stocks)-1].price <= price {
		locSpan += this.stocks[len(this.stocks)-1].span
		this.stocks = this.stocks[:len(this.stocks)-1]
	}
	this.stocks = append(this.stocks, pair{price: price, span: locSpan})
	return locSpan
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
