package models

import (
	"bytes"
	"html/template"

	"github.com/opengamer-project/og-requests/internal/models"
	"github.com/opengamer-project/og-requests/internal/store"
	"gorm.io/gorm"
)

// Claim is structure for requests representation
type Claim struct {
	gorm.Model
	UserID          uint
	User            User
	ParentCommentID *uint
	Parent          *Claim   `gorm:"foreignKey:ParentCommentID"`
	Children        []*Claim `gorm:"foreignkey:ParentCommentID"`
	RawText         string
}

func (c *Claim) Render() (template.HTML, error) {
	t, err := template.ParseFiles("./templates/components/comment.html")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, *c); err != nil {
		return "", err
	}

	return template.HTML(tpl.String()), nil

}

func NewClaim(user *models.User, text string) (*Claim, error) {
	res := &Claim{
		UserID: user.UserID,
	}
	return res, nil
}

func (c *Claim) createClaim() error {
	db := store.DB
	return db.Create(*c).Error

}
