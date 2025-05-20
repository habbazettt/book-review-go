package models

import (
	"time"
)

type Book struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Author      string    `gorm:"type:varchar(255);not null" json:"author"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	CoverImage  string    `gorm:"type:varchar(255)" json:"cover_image,omitempty"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL;" json:"category"`
	Reviews     []Review  `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;" json:"reviews,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
