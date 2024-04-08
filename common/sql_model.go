package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"colum: id"`
	CreatedAt *time.Time `json:"created_at" gorm:"colum: created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"colum: updated_at"`
}
