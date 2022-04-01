package request

import (
	validation "github.com/itgelo/ozzo-validation/v4"
	"go-klikdokter/helper/message"
)

type LoginEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginEmailRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required.Error(message.ErrReq.Message)),
		validation.Field(&req.Password, validation.Required.Error(message.ErrReq.Message)),
	)
}
