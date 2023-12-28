package service

import (
	"pig_road_server/model"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type MenuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{
		DB: db,
	}
}

func (r *MenuRepository) CreateMenu(menu *model.Menu) *model.Menu {
	result := r.DB.Create(menu)
	if result.Error != nil {
		return nil
	}
	return menu
}
