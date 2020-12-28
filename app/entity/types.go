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
	UserID         int       `json:"UserId"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	OwnedFarmIDs   []int     `json:"ownedFarmIds"`
	WorkingFarmIDs []int     `json:"workingFarmIds"`
	CreatedDate    time.Time `json:"createdDate"`
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
	FieldName   string    `json:"fieldName"`
	FarmID      int       `json:"farmId"`
	Crop        string    `json:"crop"`
	Area        float64   `json:"area"`
	CreatedDate time.Time `json:"createdDate"`
}

// Activity is an activity entity
type Activity struct {
	ActivityID   int       `json:"activityId"`
	FarmID       int       `json:"farmId"`
	Field        Field     `json:"field"`
	Tractor      Tractor   `json:"tractor"`
	User         User      `json:"user"`
	ActivityName string    `json:"activityName"`
	Area         float64   `json:"area"`
	Cost         float64   `json:"cost"`
	Revenue      float64   `json:"revenue"`
	CreatedDate  time.Time `json:"createdDate"`
}

// CostSummary is an CostSummary entity
type CostSummary struct {
	Revenue float64 `json:"revenue"`
	Cost    float64 `json:"cost"`
	Profit  float64 `json:"profit"`
}
