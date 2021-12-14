package entity

import "time"

type Photo struct {
	ID        uint      `json:"id gorm:primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Caption   string    `json:"caption" gorm:"null"`
	PhotoURL  string    `json:"photo_url" gorm:"not null"`
	UserID    int       `json:"userId"`
	User      *User     `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
