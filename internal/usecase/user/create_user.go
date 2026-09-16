package user

import (
	"context"

	"github.com/rocketseat-eduaction/gobid/internal/validator"
)

type CreateUserReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Bio      string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.Username), "username", "username is required")
	eval.CheckField(validator.NotBlank(req.Email), "email", "email is required")
	eval.CheckField(validator.NotBlank(req.Bio), "bio", "bio is required")
	eval.CheckField(validator.MinCharacters(req.Bio, 10) && validator.MaxCharacters(req.Bio, 255), "bio", "bio must be between 10 and 100 characters")
	eval.CheckField(validator.MinCharacters(req.Password, 8), "password", "password must be at least 8 characters long")
	return eval
}
