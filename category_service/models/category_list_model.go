package models

import (
	"category-service/db"
	"category-service/entities"
)

func GetListCategory() (categories []entities.Category, err error) {
	err = db.GetConn().Select(&categories, "SELECT id, name, description FROM categories")
	return
}
