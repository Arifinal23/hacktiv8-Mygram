package modelphoto

import (
	"time"
)

type Request struct {
	Title    string `json:"title" example:"A Photo"`
	Caption  string `json:"caption,omitempty" example:"Beautiful Photo"`
	PhotoURL string `json:"photo_url" example:"https://example.com/photo.jpg"`
	UserID   uint   `json:"user_id,omitempty" swaggerignore:"true"`
}

type Response struct {
	ID        uint      `json:"id,omitempty" example:"1"`
	Title     string    `json:"title"  example:"A Photo"`
	Caption   string    `json:"caption,omitempty" example:"Beautiful Photo"`
	PhotoURL  string    `json:"photo_url" example:"https://example.com/photo.jpg"`
	UserID    uint      `json:"user_id,omitempty" example:"1"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
}

type ResponseGet struct {
	ID       uint   `json:"id,omitempty" example:"1"`
	Title    string `json:"title" example:"A Photo"`
	Caption  string `json:"caption,omitempty" example:"Beautiful Photo"`
	PhotoURL string `json:"photo_url" example:"https://example.com/photo.jpg"`
	User     struct {
		Username string `json:"username" example:"jhondoe"`
		Email    string `json:"email" example:"test@example.com"`
	} `json:"user"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
}

type ResponseUpdate struct {
	ID        uint      `json:"id,omitempty"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption,omitempty"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
