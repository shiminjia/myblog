package dto

import (
	"time"
)

type Time struct {
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func NewTime() *Time {
	now := Time{
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		DeletedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	return &now
}
