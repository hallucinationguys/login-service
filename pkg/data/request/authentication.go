package request

type SignUpRequest struct {
	Name            string `valdidate:"required" json:"name"`
	Email           string `valdidate:"required" json:"email"`
	Password        string `valdidate:"required,min=8" json:"password"`
	PasswordConfirm string `valdidate:"required" json:"passwordConfirm"`
	Photo           string `valdidate:"required"`
}

type UpdateUsersRequest struct {
	Id              int    `validate:"required"`
	Name            string `validate:"required,max=200,min=2" json:"name"`
	Email           string `validate:"required,min=2,max=100" json:"email"`
	Password        string `validate:"required,min=2,max=100" json:"password"`
	PasswordConfirm string `validate:"passwordConfirm" binding:"required"`
}

type LoginRequest struct {
	Email    string `valdidate:"required, max=200,min=2" json:"email"`
	Password string `valdidate:"required, min=2,max=100" json:"password"`
}
