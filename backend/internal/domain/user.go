package domain

import "github.com/go-playground/validator/v10"

type User struct {
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"required,e164"`
	TrackId string `json:"track_id"`
}

func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)

	if err != nil {
		return ErrValidation
	}

	return nil
}
