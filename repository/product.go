package repository

import (
	"inventaris/config"
	"inventaris/model"
)

func GetAllProducts() (products []model.Product, err error) {
	db := config.DB

	if err := db.Table("produk").Scan(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(id string) (product *model.Product, err error) {

	db := config.DB

	err = db.Table("produk").Where("id = ?", id).Scan(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func AddProduct(product model.Product) error {

	db := config.DB
	err := db.Table("produk").Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(id string, updatedProduct model.Product) error {

	db := config.DB
	err := db.Table("produk").Where("id = ?", id).Updates(updatedProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id string) error {
	db := config.DB
	err := db.Table("produk").Where("id = ?", id).Delete(&model.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}
