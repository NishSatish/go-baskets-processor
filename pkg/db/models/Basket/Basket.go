package Basket

type Basket struct {
	ID       int
	stocks   []string
	name     string
	minPrice float64
}

func AddStock(stockId string) {
	sql := "INSERT INTO basket_stocks (basket_id, stock_id) VALUES (?, ?)"
}
