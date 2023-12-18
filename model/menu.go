package model

type Menu struct {
	Name     string  `json:"name"`
	Price    int     `json:"price"`
	StarRate float32 `json:"starRate"`
	Image    string  `json:"image"`
}
