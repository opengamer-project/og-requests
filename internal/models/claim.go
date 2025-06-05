package models

import (
	"bytes"
	"html/template"

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

func NewClaim(user *User, text string) (*Claim, error) {
	res := &Claim{
		UserID: (*user).ID,
	}
	return res, nil
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
