package models

import (
	"encoding/json"
	"github.com/arsalanaa44/basketball/enum"
	"time"

	"github.com/google/uuid"
)

type Basket struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	User uuid.UUID `gorm:"not null" json:"user,omitempty"`

	Data      json.RawMessage `gorm:"type:json" json:"data,omitempty"`
	Status    enum.Status     `gorm:"type:varchar(255);not null" json:"status,omitempty"`
	CreatedAt time.Time       `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time       `gorm:"not null" json:"updated_at,omitempty"`
}

type BasketRequest struct {
	Data   json.RawMessage `gorm:"type:json" json:"data,omitempty"`
	Status enum.Status     `gorm:"type:varchar(255);not null" validate:"oneof=COMPLETED PENDING" json:"status,omitempty"`
}

