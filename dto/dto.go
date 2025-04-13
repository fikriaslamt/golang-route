package dto

type UpdateStockRequest struct {
	Jumlah int `json:"jumlah" binding:"required"`
}

type UpdateStockResponse struct {
	Jumlah int `json:"jumlah"`
}
