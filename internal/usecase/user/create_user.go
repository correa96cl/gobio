package user

type CreateUserReq struct {
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required"`
	PasswordHash string `json:"password" validate:"required"`
	Bio          string `json:"bio"`
}
