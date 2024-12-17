package models

import "gorm.io/gorm"

// Comment is structure for requests representation
type Claim struct {
	gorm.Model
	UserID          uint
	User            User
	ParentCommentID *uint
	Parent          *Claim   `gorm:"foreignKey:ParentCommentID"`
	Children        []*Claim `gorm:"foreignkey:ParentCommentID"`
	RawText         string
}
