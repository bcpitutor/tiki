package models

type BannedUser struct {
	UserEmail string `json:"userEmail"`
	Details   string `json:"details"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedAt string `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
}

type BannedUserListResponse struct {
	Count       int          `json:"count"`
	Status      string       `json:"status"`
	Message     string       `json:"message"`
	BannedUsers []BannedUser `json:"data"`
}
