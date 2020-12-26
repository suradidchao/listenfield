package handler

// Response is a struct of response
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// CreateFarmPayload is a request payload for create farm endpoint
type CreateFarmPayload struct {
	FarmName    string `json:"farmName"`
	FarmOwnerID int    `json:"farmOwnerId"`
}

// AuthorizePayload is a request payload for authorize endpoint
type AuthorizePayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
