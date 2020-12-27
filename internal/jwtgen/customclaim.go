package jwtgen

// CustomClaim is a struct containing custom defined claim
type CustomClaim struct {
	Username       string `json:"username"`
	OwnedFarmIDs   []int  `json:"ownedFarmIds"`
	WorkingFarmIDs []int  `json:"workingFarmIds"`
}

// NewCustomClaim is a factory method for custom claim
func NewCustomClaim(username string, ownedFarmIDs []int, workingFarmIDs []int) CustomClaim {
	return CustomClaim{
		Username:       username,
		OwnedFarmIDs:   ownedFarmIDs,
		WorkingFarmIDs: workingFarmIDs,
	}
}
