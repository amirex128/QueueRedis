package orm

type Order struct {
	OrderId uint32 `json:"order_id" gorm:"order_id"`
	Price   uint64 `json:"price" gorm:"price"`
	Title   string `json:"title" gorm:"title"`
}
