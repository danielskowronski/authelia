package handlers

import (
	"github.com/authelia/authelia/v4/internal/authentication"
	"github.com/authelia/authelia/v4/internal/middlewares"
)

// ConfigurationBody the content returned by the configuration endpoint.
type ConfigurationBody struct {
	AvailableMethods    MethodList `json:"available_methods"`
	SecondFactorEnabled bool       `json:"second_factor_enabled"` // whether second factor is enabled or not.
	TOTPPeriod          int        `json:"totp_period"`
}

// ConfigurationGet get the configuration accessible to authenticated users.
func ConfigurationGet(ctx *middlewares.AutheliaCtx) {
	body := ConfigurationBody{}
	body.AvailableMethods = MethodList{authentication.TOTP, authentication.U2F}
	body.TOTPPeriod = ctx.Configuration.TOTP.Period

	if ctx.Configuration.DuoAPI != nil {
		body.AvailableMethods = append(body.AvailableMethods, authentication.Push)
	}

	body.SecondFactorEnabled = ctx.Providers.Authorizer.IsSecondFactorEnabled()

	ctx.Logger.Tracef("Second factor enabled: %v", body.SecondFactorEnabled)

	ctx.Logger.Tracef("Available methods are %s", body.AvailableMethods)

	err := ctx.SetJSONBody(body)
	if err != nil {
		ctx.Logger.Errorf("Unable to set configuration response in body: %s", err)
	}
}
