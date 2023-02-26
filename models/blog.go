package models

type Blog struct {
	Title  string	`json:"title,omitempty" validate:"required"`
	Content  string	`json:"content,omitempty" validate:"required"`
}