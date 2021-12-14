package modelsocialmedia

import "time"

type Request struct {
	ID             uint   `json:"id" swaggerignore:"true"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type Response struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
}

type ResponseListWrapper struct {
	SocialMedias []ResponseList `json:"social_medias"`
}
type ResponseList struct {
	ID             uint      `json:"id" example:"1"`
	CreatedAt      time.Time `json:"created_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
	Name           string    `json:"name" example:"jhondoe"`
	SocialMediaUrl string    `json:"social_media_url" example:"https://example.com/url"`
	UserID         uint      `json:"user_id" example:"1"`
	User           struct {
		ID              uint   `json:"id"  example:"1"`
		Username        string `json:"username"  example:"jhondoe"`
		ProfileImageUrl string `json:"profile_image_url" example:"https://example.com/photo.jpg"`
	} `json:"user"`
}
