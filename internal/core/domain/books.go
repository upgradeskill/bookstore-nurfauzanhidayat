package domain

import "time"

type Books struct {
	ID        int       `gorm:"autoIncrement;primaryKey" json:"id,omitempty"`
	Title     string    `json:"title"`
	Desc      string    `json:"description"`
	ISBN      string    `json:"isbn"`
	CreateAt  time.Time `json:"create_at,omitempty"`
	UpdateAt  time.Time `json:"update_at,omitempty"`
	CreatedBy int       `json:"created_by,omitempty"`
}
