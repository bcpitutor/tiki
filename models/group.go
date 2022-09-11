package models

type GroupListResponse struct {
	Count   int     `json:"count"`
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Groups  []Group `json:"data"`
}

type GroupInfoResponse struct {
	Count   int       `json:"count"`
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    GroupData `json:"data"`
}

type GroupData struct {
	AccessPerms  GroupAccessPerms `json:"accessPerms"`
	GroupName    string           `json:"groupName"`
	CreatedAt    int64            `json:"createdAt"`
	CreatedBy    string           `json:"createdBy"`
	UpdatedAt    int64            `json:"updatedAt"`
	UpdatedBy    string           `json:"updatedBy"`
	GroupMembers []string         `json:"groupMembers"`
}

type GroupAccessPerm_Domain struct {
	Create bool `json:"create"`
	Delete bool `json:"delete"`
}

type GroupAccessPerm_Group struct {
	Create    bool `json:"create"`
	Delete    bool `json:"delete"`
	AddMember bool `json:"addMember"`
	DelMember bool `json:"delMember"`
}
type GroupAccessPerm_Ticket struct {
	Create bool `json:"create"`
	Delete bool `json:"delete"`
	Show   bool `json:"show"`
}

type GroupAccessPerms struct {
	Domain GroupAccessPerm_Domain `json:"domain"`
	Group  GroupAccessPerm_Group  `json:"group"`
	Ticket GroupAccessPerm_Ticket `json:"ticket"`
}
type Group struct {
	GroupName    string   `json:"groupName"`
	CreatedAt    int64    `json:"createdAt"`
	CreatedBy    string   `json:"createdBy"`
	UpdatedAt    int64    `json:"updatedAt"`
	UpdatedBy    string   `json:"updatedBy"`
	GroupMembers []string `json:"groupMembers"`
}
