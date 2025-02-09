package main

import (
	"github.com/abyanmajid/matcha"
	"github.com/abyanmajid/matcha/email"
	"github.com/abyanmajid/matcha/logger"
	"github.com/abyanmajid/matcha/openapi"
	"github.com/abyanmajid/matcha/reference"
	"github.com/abyanmajid/thorfinn/internal"
	"github.com/abyanmajid/thorfinn/internal/api"
)

const (
	AuthRegisterPath              = "/auth/register"
	AuthVerifyEmailPath           = "/auth/verify-email"
	AuthLoginPath                 = "/auth/login"
	AuthLogoutPath                = "/auth/logout"
	AuthSendEmailVerificationPath = "/auth/send-email-verification"
	AuthSendPasswordResetLinkPath = "/auth/send-password-reset-link"
	AuthResetPasswordPath         = "/auth/reset-password"
	AuthOtpSendPath               = "/auth/otp/send"
	AuthOtpVerifyPath             = "/auth/otp/verify"

	UsersGetAllPath = "/users"
	UsersGetPath    = "/users/{id}"
	UsersUpdatePath = "/users/{id}"
	UsersDeletePath = "/users/{id}"
)

func main() {
	app := matcha.New()

	isDev, config := internal.ConfigureEnv()

	queries, err := internal.CreateQueryClient(config.DatabaseUrl)
	if err != nil {
		logger.Fatal("Failed to create query client: %v", err)
	}

	mailer := email.NewClient(email.Config{
		Host:     config.SmtpHost,
		Port:     config.SmtpPort,
		Username: config.SmtpUser,
		Password: config.SmtpPassword,
	}, "templates")

	resources, err := api.CreateApiResources(&api.Utils{
		IsDev:   &isDev,
		Config:  config,
		Queries: queries,
		Mailer:  mailer,
	})
	if err != nil {
		logger.Fatal("Failed to create resources: %v", err)
	}

	app.Documentation("/docs", openapi.Meta{
		OpenAPI:        "3.0.0",
		PackageName:    "Thorfinn API",
		PackageVersion: "0.1.0",
	})

	// Authentication resources
	app.Post(AuthRegisterPath, resources.AuthResources.Register)
	app.Put(AuthVerifyEmailPath, resources.AuthResources.VerifyEmail)
	app.Post(AuthLoginPath, resources.AuthResources.Login)
	app.Post(AuthLogoutPath, resources.AuthResources.Logout)
	app.Post(AuthSendEmailVerificationPath, resources.AuthResources.SendEmailVerification)
	app.Post(AuthSendPasswordResetLinkPath, resources.AuthResources.SendPasswordResetLink)
	app.Put(AuthResetPasswordPath, resources.AuthResources.ResetPassword)
	app.Post(AuthOtpSendPath, resources.AuthResources.OtpSend)
	app.Post(AuthOtpVerifyPath, resources.AuthResources.OtpVerify)

	// User management resources
	app.Get(UsersGetAllPath, resources.UsersResources.GetAllUsers)
	app.Get(UsersGetPath, resources.UsersResources.GetUser)
	app.Put(UsersUpdatePath, resources.UsersResources.UpdateUser)
	app.Delete(UsersDeletePath, resources.UsersResources.DeleteUser)

	app.Reference("/reference", &reference.Options{
		Source: "/docs",
	})

	app.Serve(":8080")
}
