package service

import (
	"pig_road_server/model"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type RestaurantRepository struct {
	DB *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *RestaurantRepository {
	return &RestaurantRepository{
		DB: db,
	}
}

func (r *RestaurantRepository) CreateRestaurant(name, address string, rating, lat, lon float32) *model.Restaurant {
	restaurant := model.Restaurant{
		Name:      name,
		Address:   address,
		Rating:    rating,
		Latitude:  lat,
		Longitude: lon,
	}
	result := r.DB.Create(&restaurant)
	if result.Error != nil {
		return nil
	}
	return &restaurant
}
