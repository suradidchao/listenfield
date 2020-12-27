package farmworker

// IRepo is an interface of farmworker repository
type IRepo interface {
	Create(farmID int, userID int) (fwID int, err error)
}

// Repo is an implementation of farworker repository
type Repo struct {
	farmWorkerAdapter IAdapter
}

// Create is a method for inserting user id and farm id into farmworker
func (r Repo) Create(farmID int, userID int) (fwID int, err error) {
	return r.farmWorkerAdapter.Create(farmID, userID)
}

// NewRepo is a factory method of farm worker repository
func NewRepo(fwa IAdapter) Repo {
	return Repo{
		farmWorkerAdapter: fwa,
	}
}
