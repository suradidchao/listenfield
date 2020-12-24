package repo

import "time"

// Farm is a farm entity
type Farm struct {
	FarmID      int       `json:"farm_id"`
	FarmName    string    `json:"farm_name"`
	FarmOwner   Farmer    `json:"farm_owner"`
	Fields      []Field   `json:"fields"`
	Tractors    []Tractor `json:"tractors"`
	FarmWorkers []Farmer  `json:"farm_workers"`
	CreatedDate time.Time `json:"created_date"`
}

// Farmer is a farm entity
type Farmer struct {
	FarmerID    int       `json:"farmer_id"`
	FarmerName  string    `json:"farmer_name"`
	CreatedDate time.Time `json:"created_date"`
}

// Tractor is a farm entity
type Tractor struct {
	TractorID   int       `json:"tractor_id"`
	TractorName string    `json:"tractor_name"`
	FarmID      int       `json:"farm_id"`
	CreatedDate time.Time `json:"created_date"`
}

// Field is a farm entity
type Field struct {
	FieldID     int       `json:"field_id"`
	Crop        string    `json:"crop"`
	Status      string    `json:"status"`
	CreatedDate time.Time `json:"created_date"`
}

// Activity is a farm entity
type Activity struct {
	ActivityID   int       `json:"activity_id"`
	FarmID       int       `json:"farm_id"`
	FieldID      int       `json:"field_id"`
	ActivityName string    `json:"activity_name"`
	Cost         float64   `json:"cost"`
	CreatedDate  time.Time `json:"created_date"`
}
