package model

type Product struct {
	ID          int     `json:"id" gorm:"primaryKey;column:id" query:"id"`
	Name        string  `json:"nama" gorm:"column:nama" query:"nama"`
	Description string  `json:"deskripsi" gorm:"column:deskripsi" query:"deskripsi"`
	Price       float64 `json:"harga" gorm:"column:harga" query:"harga"`
	Category    string  `json:"kategori" gorm:"column:kategori" query:"kategori"`
	Filepath    string  `json:"filepath" gorm:"column:filepath" query:"filepath"`
}
type Inventory struct {
	ID        int    `json:"id" gorm:"primaryKey;column:id" query:"id"`
	ProductID int    `json:"produk_id" gorm:"column:produk_id" query:"produk_id"`
	Quantity  int    `json:"jumlah" gorm:"column:jumlah" query:"jumlah"`
	Location  string `json:"lokasi" gorm:"column:lokasi" query:"lokasi"`
}

type Order struct {
	ID        int    `json:"id" gorm:"primaryKey;column:id" query:"id"`
	ProductID int    `json:"produk_id" gorm:"column:produk_id" query:"produk_id"`
	Quantity  int    `json:"jumlah" gorm:"column:jumlah" query:"jumlah"`
	Date      string `json:"tanggal" gorm:"column:tanggal" query:"tanggal"`
}
