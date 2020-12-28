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

// UserCreatePayload is a request paylod for create user endpoint
type UserCreatePayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// AddFarmWorkerPayload is a request payload for adding worker to a farm endpoint
type AddFarmWorkerPayload struct {
	WorkerID int `json:"workerId"`
}

// AddTractorPayload is a request payload for adding worker to a farm endpoint
type AddTractorPayload struct {
	TractorName string `json:"tractorName"`
}

// AddFieldPayload is a request payload for adding field to a farm endpoint
type AddFieldPayload struct {
	FieldName string  `json:"fieldName"`
	Crop      string  `json:"crop"`
	Status    string  `json:"status"`
	Area      float64 `json:"area"`
}
