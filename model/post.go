package model

import "time"

type Post struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Subject   string    `json:"subject"`
	Content   string    `json:"content"`
}

func (o Post) TableName() string {
	return "post"
}
