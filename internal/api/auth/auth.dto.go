package auth_features

type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type ConfirmEmailRequest struct {
	Token string `json:"token"`
}

type ConfirmEmailResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message       string `json:"message"`
	AccessToken   string `json:"access_token"`
	RefreshTokens string `json:"refresh_token"`
}

type LogoutRequest struct{}

type LogoutResponse struct {
	Message string `json:"message"`
}

type SendVerificationEmailRequest struct {
	Email string `json:"email"`
}

type SendVerificationEmailResponse struct {
	Message string `json:"message"`
}

type SendPasswordResetRequest struct {
	Email string `json:"email"`
}

type SendPasswordResetResponse struct {
	Message string `json:"message"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
}

type OtpSendRequest struct {
	Email string `json:"email"`
}

type OtpSendResponse struct {
	Message   string `json:"message"`
	OtpCodeId string `json:"otp_code_id"`
}

type OtpVerifyRequest struct {
	OtpCodeId string `json:"otp_code_id"`
	OtpCode   string `json:"otp_code"`
}

type OtpVerifyResponse struct {
	Message string `json:"message"`
}
