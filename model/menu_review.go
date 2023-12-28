package model

import "time"

type MenuReview struct {
	Index   int
	Rating  int
	Comment string
	Author  string
	Date    time.Time
	MenuID  int
}
