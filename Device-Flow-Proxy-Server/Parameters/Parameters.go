package Parameters

type ParamAuthResponse struct {
	DeviceCode      string `json:"device_code"  validate:"required"`
	UserCode 	string `json:"user_code" validate:"required"`
	VerificationURI string `json:"verification_uri" validate:"required"`
}

type ParamTokenResponse struct {
	AccessToken    string `json:"access_token"  validate:"required"`
	TokenType      string `json:"token_type" validate:"required"`
	ExpiresInToken string `json:"expires_in" validate:"required"`
}

