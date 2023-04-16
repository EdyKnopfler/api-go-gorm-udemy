package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name  string
	Email string

	// Auto-relacionamento muitos-para-muitos
	Connections []Member `gorm:"many2many:member_connections"`

	// Um-para-muitos
	Notes []Note `gorm:"foreignKey:MemberID"`
}

type Note struct {
	gorm.Model

	// Possui-um
	Author Member `gorm:"foreignKey:ID"`

	Text     string
	MemberID string
}
