package models

type Session struct {
	SessionId       string `json:"sessionId"`
	SessionOwner    string `json:"sessionOwner"`   //  email
	SessionDetails  string `json:"sessionDetails"` // comment
	ExpiresAt       string `json:"expiresAt"`
	Epoch           int64  `json:"epoch"`
	SessionExpEpoch int64  `json:"sessionExpEpoch"`
	RefreshCount    int    `json:"refreshCount"`
	Revoked         bool   `json:"revoked"`
}

type SessionListResponse struct {
	Count    int       `json:"count"`
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Sessions []Session `json:"data"`
}
