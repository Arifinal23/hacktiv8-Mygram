package entity

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `json:"message" gorm:"not null"`
	User      *User     `gorm:"foreignKey:UserID"`
	Photo     *Photo    `gorm:"foreignKey:PhotoID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
