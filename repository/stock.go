package repository

import (
	"errors"
	"inventaris/config"
	"inventaris/dto"
	"inventaris/model"
)

func GetStock(productID string) (int, error) {

	db := config.DB
	data := model.Inventory{}
	err := db.Table("produk").Where("id = ?", productID).Scan(&data).Error
	if err != nil {
		return 0, err
	}
	if data.Quantity == 0 {
		return 0, errors.New("stock is empty")
	}
	return data.Quantity, nil
}

func UpdateStock(productID string, payload dto.UpdateStockRequest) (inventory model.Inventory, err error) {
	db := config.DB
	data := model.Inventory{}
	err = db.Table("produk").Where("id = ?", productID).Scan(&data).Error
	if err != nil {
		return model.Inventory{}, err
	}
	if data.Quantity == 0 {
		return model.Inventory{}, errors.New("stock is empty")
	}

	data.Quantity = payload.Jumlah

	err = db.Table("produk").Where("id = ?", productID).Updates(data).Scan(&inventory).Error
	if err != nil {
		return model.Inventory{}, err
	}
	return inventory, nil
}
