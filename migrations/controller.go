package migrations

import "github.com/go-pg/pg/v10"

// ControllerDeps describes the necessary dependencies for the Migration controller
type ControllerDeps struct {
	Service  *Service
	database *pg.DB
}

// Controller governs the migrating up, down, and initializing of migrations
type Controller struct {
	Service  *Service
	database *pg.DB
}

// CreateControllerDeps helps inject dependencies into a controller
func CreateControllerDeps(db *pg.DB) ControllerDeps {
	svc := &Service{database: db}
	return ControllerDeps{Service: svc, database: db}
}

// NewController initializes a new controller using the dependencies given
// to it
func NewController(deps ControllerDeps) *Controller {
	return &Controller{database: deps.database}
}
