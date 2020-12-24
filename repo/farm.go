package repo

// IFarmRepo is an interface for farm repository
type IFarmRepo interface {
	CreateFarm(farm Farm) (int, error)
}
