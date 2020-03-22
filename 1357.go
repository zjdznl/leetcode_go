package leetcode
//
//type Cashier struct {
//	N            int
//	Discount     int
//	ProductPrice map[int]int
//
//	CustomerNum int
//}
//
//func Constructor(n int, discount int, products []int, prices []int) Cashier {
//	c := Cashier{
//		N:            n,
//		Discount:     discount,
//		ProductPrice: make(map[int]int, len(products)),
//	}
//	for i, product := range products {
//		c.ProductPrice[product] = prices[i]
//	}
//	return c
//}
//
//func (this *Cashier) GetBill(product []int, amount []int) float64 {
//	totalPrice := 0.0
//	for i, p := range product {
//		totalPrice += float64(amount[i] * this.ProductPrice[p])
//	}
//	this.CustomerNum++
//	if this.CustomerNum%this.N == 0 {
//		totalPrice -= totalPrice * float64(this.Discount) / float64(100)
//	}
//	return float64(totalPrice)
//}

/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */
