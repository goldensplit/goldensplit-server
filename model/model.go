// Package model provides TODO: Docs.
package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Model TODO: Docs.
type Model interface {
	mock()
}

// User TODO: Docs.
type User struct {
	gorm.Model
	Name  string
	Email string
	Info  string
}

// Session TODO: Docs.
type Session struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Video TODO: Docs.
type Video struct {
	gorm.Model
	UploaderID uint
	Info       string
	Tags       []string `gorm:"type:varchar(255)[]"`
}

func (User) mock()    {}
func (Session) mock() {}
func (Video) mock()   {}
