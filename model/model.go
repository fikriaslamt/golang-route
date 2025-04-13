package model

type Product struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Name        string  `json:"nama" gorm:"size:255;not null"`
	Description string  `json:"deskripsi" gorm:"type:text"`
	Price       float64 `json:"harga" gorm:"type:decimal(10,2);not null"`
	Category    string  `json:"kategori" gorm:"size:100;not null"`
}

type Inventory struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	ProductID int    `json:"produk_id" gorm:"not null"`
	Quantity  int    `json:"jumlah" gorm:"not null"`
	Location  string `json:"lokasi" gorm:"size:255;not null"`
}

type Order struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	ProductID int    `json:"produk_id" gorm:"not null"`
	Quantity  int    `json:"jumlah" gorm:"not null"`
	Date      string `json:"tanggal" gorm:"type:date;not null"`
}
