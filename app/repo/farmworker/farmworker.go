package farmworker

// IRepo is an interface of farmworker repository
type IRepo interface {
	Create(farmID int, userID int) (fwID int, err error)
	GetAllByFarmID(farmID int) (userIDs []int, err error)
	Delete(farmID int, userID int) (err error)
}

// Repo is an implementation of farworker repository
type Repo struct {
	farmWorkerAdapter IAdapter
}

// Create is a method for inserting user id and farm id into farmworker
func (r Repo) Create(farmID int, userID int) (fwID int, err error) {
	return r.farmWorkerAdapter.Create(farmID, userID)
}

// GetAllByFarmID is a method for get all workers id from a farm
func (r Repo) GetAllByFarmID(farmID int) (userIDs []int, err error) {
	return r.farmWorkerAdapter.GetAllByFarmID(farmID)
}

// Delete is a method for deleting farm worker from a farm
func (r Repo) Delete(farmID int, userID int) (err error) {
	return r.farmWorkerAdapter.Delete(farmID, userID)
}

// NewRepo is a factory method of farm worker repository
func NewRepo(fwa IAdapter) Repo {
	return Repo{
		farmWorkerAdapter: fwa,
	}
}
