package main

import (
	"log"
	"pig_road_server/model"
	"pig_road_server/service"
)

func main() {
	ds, err := service.NewDatabaseService()
	if err != nil {
		log.Fatal(err)
	}
	// restaurantService := service.NewRestaurantRepository(ds.GetDB())
	// result := restaurantService.CreateRestaurant("파이브가이즈 강남점", "서울특별시 서초구 강남대로 435", 4.0, 37.5012, 127.0257)
	// log.Println(result)

	menus := model.MockingMenus
	menuService := service.NewMenuRepository(ds.GetDB())
	for _, menu := range menus {
		result := menuService.CreateMenu(&menu)
		log.Println(result)
	}
}
