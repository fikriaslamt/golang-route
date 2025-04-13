package repository

import (
	"inventaris/config"
	"inventaris/model"
)

func GetOrderByID(id string) (order *model.Order, err error) {
	db := config.DB

	err = db.Table("order").Where("id = ?", id).Scan(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil

}

func CreateOrder(order model.Order) error {

	db := config.DB
	err := db.Table("order").Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}
