package model

type CreateAppRequest struct {
	Name string `json:"name"`
}

type UpdateAppRequest struct {
	Name string `json:"name"`
}

type UserQuery struct {
	ID      string `query:"id"`
	AppID   string `query:"app_id"`
	Keyword string `query:"keyword"`
	OrderBy string `query:"order_by"`
	SortBy  string `query:"sort_by"`
	Limit   int    `query:"limit"`
	Offser  int    `query:"offset"`
}

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Ticket   string `json:"ticket"`
	Password string `json:"password"`
}

type RequestVerifyEmailRequest struct {
	Email string `json:"email"`
}

type VerifyEmailRequest struct {
	Ticket string `json:"ticket"`
	Email  string `json:"email"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
