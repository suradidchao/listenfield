package entity

import "time"

// Farm is a farm entity
type Farm struct {
	FarmID      int       `json:"farmId"`
	FarmName    string    `json:"farmName"`
	FarmOwner   User      `json:"farmOwner"`
	Fields      []Field   `json:"fields"`
	Tractors    []Tractor `json:"tractors"`
	FarmWorkers []User    `json:"farmWorkers"`
	CreatedDate time.Time `json:"createdDate"`
}

// User is a farm entity
type User struct {
	UserID      int       `json:"UserId"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	CreatedDate time.Time `json:"createdDate"`
}

// Tractor is a farm entity
type Tractor struct {
	TractorID   int       `json:"tractorId"`
	TractorName string    `json:"tractorName"`
	FarmID      int       `json:"farmId"`
	CreatedDate time.Time `json:"createdDate"`
}

// Field is a farm entity
type Field struct {
	FieldID     int       `json:"fieldId"`
	Crop        string    `json:"crop"`
	Status      string    `json:"status"`
	CreatedDate time.Time `json:"createdDate"`
}

// Activity is a farm entity
type Activity struct {
	ActivityID   int       `json:"activityId"`
	FarmID       int       `json:"farmId"`
	FieldID      int       `json:"fieldId"`
	ActivityName string    `json:"activityName"`
	Cost         float64   `json:"cost"`
	CreatedDate  time.Time `json:"createdDate"`
}
