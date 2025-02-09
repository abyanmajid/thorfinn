package auth_features

import "github.com/abyanmajid/matcha/openapi"

type DerivedAuthResources struct {
	Register              *openapi.Resource
	VerifyEmail           *openapi.Resource
	Login                 *openapi.Resource
	Logout                *openapi.Resource
	SendEmailVerification *openapi.Resource
	SendPasswordResetLink *openapi.Resource
	ResetPassword         *openapi.Resource
	OtpSend               *openapi.Resource
	OtpVerify             *openapi.Resource
}

func Derive(handlers *AuthHandlers) (*DerivedAuthResources, error) {
	authResources := NewAuthResources(handlers)
	registerResource, err := authResources.RegisterResource()
	if err != nil {
		return nil, err
	}

	confirmEmailResource, err := authResources.ConfirmEmailResource()
	if err != nil {
		return nil, err
	}

	loginResource, err := authResources.LoginResource()
	if err != nil {
		return nil, err
	}

	logoutResource, err := authResources.LogoutResource()
	if err != nil {
		return nil, err
	}

	sendEmailVerificationResource, err := authResources.SendEmailVerificationResource()
	if err != nil {
		return nil, err
	}

	sendPasswordResetLinkResource, err := authResources.SendPasswordResetLinkResource()
	if err != nil {
		return nil, err
	}

	resetPasswordResource, err := authResources.ResetPasswordResource()
	if err != nil {
		return nil, err
	}

	otpSendResource, err := authResources.OtpSendResource()
	if err != nil {
		return nil, err
	}

	otpVerifyResource, err := authResources.OtpVerifyResource()
	if err != nil {
		return nil, err
	}

	return &DerivedAuthResources{
		Register:              registerResource,
		VerifyEmail:           confirmEmailResource,
		Login:                 loginResource,
		Logout:                logoutResource,
		SendEmailVerification: sendEmailVerificationResource,
		SendPasswordResetLink: sendPasswordResetLinkResource,
		ResetPassword:         resetPasswordResource,
		OtpSend:               otpSendResource,
		OtpVerify:             otpVerifyResource,
	}, nil
}
