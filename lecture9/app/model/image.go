package model

type Image struct {
	Image_id    int    `json:"image_id"`
	Image_link  string `json:"image_link"`
	Description string `json:"description"`
	UserID      int64  `json:"user_id"`
}
