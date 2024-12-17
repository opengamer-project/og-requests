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

func (c *Claim) Render() (template.HTML, error) {
	t, err := template.ParseFiles("./templates/components/comment.html")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		return "", err
	}

	return template.HTML(tpl.String()), nil

}
