package model

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	ID	   uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string  `json:"name"`
	Price    int     `json:"price"`
	StarRate float32 `json:"starRate"`
	Image    string  `json:"image"`
	RestaurantID uint `json:"restaurantId"`
}

var MockingMenus = []Menu{
	Menu {
		Name: "햄버거",
		Price: 13400,
		StarRate: 4.5,
		Image: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.istockphoto.com%2Fvector%2Fburger-ico",
		RestaurantID: 1,
	},
	Menu {
		Name: "치즈 버거",
		Price: 14900,
		StarRate: 4.5,
		Image: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.istockphoto.com%2Fvector%2Fburger-ico",
		RestaurantID: 1,
	},
	Menu {
		Name: "베이컨 버거",
		Price: 15900,
		StarRate: 4.5,
		Image: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.istockphoto.com%2Fvector%2Fburger-ico",
		RestaurantID: 1,
	},
	Menu {
		Name: "베이컨 치즈 버거",
		Price: 17400,
		StarRate: 4.5,
		Image: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.istockphoto.com%2Fvector%2Fburger-ico",
		RestaurantID: 1,
	},
	Menu {
		Name: "리틀 햄버거",
		Price: 9900,
		StarRate: 4.5,
		Image: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.istockphoto.com%2Fvector%2Fburger-ico",
		RestaurantID: 1,
	},
	Menu {
		Name: "리틀 치즈 버거",
		Price: 11400,
		StarRate: 4.5,
		Image: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.istockphoto.com%2Fvector%2Fburger-ico",
		RestaurantID: 1,
	},
	Menu {
		Name: "리틀 베이컨 버거",
		Price: 12400,
		StarRate: 4.5,
		Image: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.istockphoto.com%2Fvector%2Fburger-ico",
		RestaurantID: 1,
	},
	Menu {
		Name: "리틀 베이컨 치즈 버거",
		Price: 13900,
		StarRate: 4.5,
		Image: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.istockphoto.com%2Fvector%2Fburger-ico",
		RestaurantID: 1,
	},
	Menu {
		Price: 9900,
		StarRate: 4.5,
		Image: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.istockphoto.com%2Fvector%2Fburger-ico",
		RestaurantID: 1,
	},
	Menu {
		Name: "치즈 베지 샌드위치",
		Price: 11400,
		StarRate: 4.5,
		Image: "https://www.google.com/url?sa=i&url=http%3A%2F%2Fwww.istockphoto.com%2Fvector%2Fchicken-sandwich-icon-gm1183950949-332843447&psig=AOvVaw2Q4Z4X6Z4Z4Z4Z4Z4Z4Z4Z&ust=1629789850000000&source=images&cd=vfe&ved=0CAsQjRxqFwoTCJjQ9ZqHgvICFQAAAAAdAAAAABAD",
		RestaurantID: 1,
	},
	Menu {
		Name: "그릴드 치즈",
		Price: 8400,
		StarRate: 4.5,
		Image: "http://www.istockphoto.com/vector/chicken-sandwich-icon-gm1183950949-332843447",
		RestaurantID: 1,
	},
	Menu {
		Name: "BLT 샌드위치",
		Price: 10900,
		StarRate: 4.5,
		Image: "http://www.istockphoto.com/vector/chicken-sandwich-icon-gm1183950949-332843447",
		RestaurantID: 1,
	},
}