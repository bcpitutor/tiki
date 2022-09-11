package models

type AwsCredentials struct {
	AccessKeyId     string `json:"aws_access_key_id"`
	SecretAccessKey string `json:"aws_secret_access_key"`
	SessionToken    string `json:"aws_session_token"`
	Region          string `json:"aws_region"`
}

type AwsConsoleCredentials struct {
	SessionId    string `json:"sessionId"`
	SessionKey   string `json:"sessionKey"`
	SessionToken string `json:"sessionToken"`
	Region       string `json:"Region"`
}

type AwsFederationResponse struct {
	SigninToken string `json:"SigninToken"`
}
