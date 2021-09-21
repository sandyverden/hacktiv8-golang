package structs

type Orders struct {
	CustomerName string  `json:"customerName" gorm:"column:customer_name"`
	OrderAt      string  `json:"orderedAt" gorm:"column:ordered_at"`
	Items        []Items `json:"items" gorm:"embedded"`
}

type Items struct {
	ItemCode    string `json:"itemCode" gorm:"column:item_code"`
	Description string `json:"description" gorm:"column:description"`
	Quantity    int    `json:"quantity" gorm:"column:quantity"`
}
