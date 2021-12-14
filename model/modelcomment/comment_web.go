package modelcomment

import "time"

type Request struct {
	Message string `json:"message" example:"A Photo"`
	PhotoID uint   `json:"photo_id,omitempty" example:"1"`
	UserID  uint   `json:"user_id,omitempty" swaggerignore:"true"`
}

type RequestUpdate struct {
	Message string `json:"message" example:"A Photo"`
	UserID  uint   `json:"user_id,omitempty" swaggerignore:"true"`
}

type ResponseInsert struct {
	ID        uint      `json:"id" example:"1"`
	Message   string    `json:"message" example:"A Photo"`
	PhotoID   uint      `json:"photo_id" example:"1"`
	UserID    uint      `json:"user_id" example:"1"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
}

type ResponseUpdate struct {
	ID        uint      `json:"id" example:"1"`
	Message   string    `json:"message" example:"A Photo"`
	PhotoID   uint      `json:"photo_id" example:"1"`
	UserID    uint      `json:"user_id" example:"1"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-11-03T01:52:41.035Z"`
}

type Response struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	User      struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"user"`
	Photo struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption,omitempty"`
		PhotoURL string `json:"photo_url"`
		UserID   uint   `json:"user_id"`
	} `json:"photo"`
}
