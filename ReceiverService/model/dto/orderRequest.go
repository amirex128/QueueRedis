package dto

// OrderRequest is a struct that represents the request body for the order
type OrderRequest struct {
	OrderId uint32 `json:"order_id" binding:"required"`
	Price   uint64 `json:"price" binding:"required"`
	Title   string `json:"title" binding:"required"`
}
