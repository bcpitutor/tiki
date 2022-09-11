package models

type DomainQueryResponse struct {
	Count   int      `json:"count"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Domains []Domain `json:"domains"`
	Data    Domain   `json:"data"`
}

type Domain struct {
	DomainPath    string `json:"domainPath"`
	CreatedAt     int64  `json:"createdAt"`
	CreatedBy     string `json:"createdBy"`
	UpdatedAt     int64  `json:"updatedAt"`
	UpdatedBy     string `json:"updatedBy"`
	Parent        string `json:"parent"`
	OwnerGroup    string `json:"ownerGroup"`
	DomainComment string `json:"domainComment"`
}
