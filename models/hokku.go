package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Hokku struct {
	Id      int       `json:"id" form:"id"`
	Title   string    `json:"title" form:"title"`
	Content string    `json:"content" form:"content"`
	Created time.Time `json:"created" form:"created"`
	OwnerId int       `json:"ownerId" form:"ownerId"`
	ThemeId int       `json:"themeId" form:"themeId"`
}

func (h *Hokku) Validate() error {
	return validation.ValidateStruct(
		h,
		validation.Field(&h.Title, validation.Required, validation.Length(1, 255)),
		validation.Field(&h.Content, validation.Required),
	)
}
