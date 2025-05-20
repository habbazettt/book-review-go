package models

import (
	"time"
)

type User struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	Email         string         `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password      string         `gorm:"not null" json:"-"`
	Role          string         `gorm:"type:varchar(50);default:user" json:"role"`
	Reviews       []Review       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"reviews,omitempty"`
	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"-"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}
