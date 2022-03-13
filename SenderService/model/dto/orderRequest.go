package dto

type OrderRequest struct {
	OrderId uint32 `json:"order_id"`
	Price   uint64 `json:"price"`
	Title   string `json:"title"`
}
