package models

type TicketQueryResponse struct {
	Count   int      `json:"count"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Tickets []Ticket `json:"tickets"`
	Data    Ticket   `json:"data"`
}

type AWSTicketAssumeRoleDetails struct {
	RoleArn string `json:"roleArn"`
	TTL     int64  `json:"ttl"`
}

type AWSTicketAwsPermissions struct {
	Resource string   `json:"resource"`
	Effect   string   `json:"effect"`
	Action   []string `json:"action"`
}

type Ticket struct {
	TicketPath   string                     `json:"ticketPath"`
	TicketType   string                     `json:"ticketType"`
	TicketRegion string                     `json:"ticketRegion"`
	TicketInfo   string                     `json:"ticketInfo"`
	ATSD         AWSTicketAssumeRoleDetails `json:"assumeRoleDetails"`
	ATAP         AWSTicketAwsPermissions    `json:"awsPermissions"`
	CreatedAt    string                     `json:"createdAt"`
	CreatedBy    string                     `json:"createdBy"`
	UpdatedAt    string                     `json:"updatedAt"`
	UpdatedBy    string                     `json:"updatedBy"`
	OwnerGroups  []string                   `json:"ownersGroup"` // TODO: check field name
}
