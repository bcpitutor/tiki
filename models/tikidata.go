package models

type TikiData struct {
	GoogleJwtToken string `json:"googleJwtToken"`
	IdentityType   string `json:"identityType"`
	MachineSN      string `json:"machineSerialNumber"`
	HashValue      string `json:"hashValue"`
}

var (
	TikiInfo *TikiData
)

func (m *TikiData) GetGoogleToken() string {
	return m.GoogleJwtToken
}

func (m *TikiData) SetGoogleJwtToken(googleJwtToken string) {
	m.GoogleJwtToken = googleJwtToken
}

func (m *TikiData) GetIdentityType() string {
	return m.IdentityType
}

func (m *TikiData) SetIdentityType(identityType string) {
	m.IdentityType = identityType
}

func (m *TikiData) GetMachineSN() string {
	return m.MachineSN
}

func (m *TikiData) SetMachineSN(machineSN string) {
	m.MachineSN = machineSN
}
