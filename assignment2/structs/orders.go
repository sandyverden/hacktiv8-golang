package structs

type Orders struct {
	orders       string
	order_id     int
	customerName string
	ordered_at   string
}

type OrdersBy struct {
	items       string
	itemID      int
	itemCode    string
	description string
	quantity    int
	order_id    int
}
