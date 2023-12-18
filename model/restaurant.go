package model

type Restaurant struct {
	Name       string     `json:"name"`
	Coordinate Coordinate `json:"coordinate"`
	ImageURL   string     `json:"imageURL"`
	Menus      []Menu     `json:"menus"`
}

func GetKoreanRestaurants() []Restaurant {
	restaurants := []Restaurant{
		Restaurant{
			Name: "김밥천국",
			Coordinate: Coordinate{
				Latitude:  37.5665,
				Longitude: 126.9780,
			},
			ImageURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Seoul-Cityscape_Namsan_Night.jpg/1200px-Seoul-Cityscape_Namsan_Night.jpg",
			Menus: []Menu{
				Menu{
					Name:     "김밥",
					Price:    3000,
					StarRate: 4.0,
					Image:    "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Seoul-Cityscape_Namsan_Night.jpg/1200px-Seoul-Cityscape_Namsan_Night.jpg",
				},
			},
		},
		Restaurant{
			Name: "돈까스 천국",
			Coordinate: Coordinate{
				Latitude:  37.5665,
				Longitude: 126.9780,
			},
			ImageURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Seoul-Cityscape_Namsan_Night.jpg/1200px-Seoul-Cityscape_Namsan_Night.jpg",
			Menus: []Menu{
				Menu{
					Name:     "돈까스",
					Price:    5000,
					StarRate: 5.0,
					Image:    "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Seoul-Cityscape_Namsan_Night.jpg/1200px-Seoul-Cityscape_Namsan_Night.jpg",
				},
			},
		},
		Restaurant{
			Name: "초밥 천국",
			Coordinate: Coordinate{
				Latitude:  37.5665,
				Longitude: 126.9780,
			},
			ImageURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Seoul-Cityscape_Namsan_Night.jpg/1200px-Seoul-Cityscape_Namsan_Night.jpg",
			Menus: []Menu{
				Menu{
					Name:     "초밥",
					Price:    10000,
					StarRate: 3.0,
					Image:    "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4e/Seoul-Cityscape_Namsan_Night.jpg/1200px-Seoul-Cityscape_Namsan_Night.jpg",
				},
			},
		},
	}

	return restaurants
}
