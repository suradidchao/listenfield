package handler

// Response is a struct of response
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// CreateFarmPayload is a request payload for create farm endpoint
type CreateFarmPayload struct {
	FarmName    string `json:"farm_name"`
	FarmOwnerID int    `json:"farm_owner"`
}
