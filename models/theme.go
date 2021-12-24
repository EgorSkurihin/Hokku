package models

import validation "github.com/go-ozzo/ozzo-validation"

type Theme struct {
	Id    int    `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
}

func (t *Theme) Validate() error {
	return validation.ValidateStruct(
		t,
		validation.Field(&t.Title, validation.Required, validation.Length(1, 40)),
	)
}
