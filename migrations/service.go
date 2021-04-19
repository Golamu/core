package migrations

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

// Service is the service that you use to do things like run migrations,
// initialize them, etc
type Service struct {
	directory string
	database  *pg.DB
}

// NewService creates a new service with the database connection provided
func NewService(db *pg.DB) *Service {
	return &Service{database: db}
}

// Initialize checks if gopg has been initialized, and if it has not, it
// will run the gopg migrations. This will then return if the initialize
// has actually run the migration init
func (svc *Service) Initialize() (bool, error) {
	find := &pgClass{Name: "gopg_migrations", Kind: 'r'}
	count, err := svc.database.Model(find).SelectAndCount()
	if err != nil {
		return false, err
	}

	if count > 0 {
		return false, nil
	}

	_, _, err = migrations.Run(svc.database, "init")
	if err != nil {
		return false, err
	}

	return true, nil
}

// SetDirectory sets the find directory for automatically running and sql migrations.
// this MUST be the fully qualified path with the `.sql` files inside.
func (svc *Service) SetDirectory(arg string) {
	svc.directory = arg
	migrations.DefaultCollection.DiscoverSQLMigrations(arg)
}

// Up runs all migrations up to the latest. Returns old version, new version, and the
// error if there was any
func (svc *Service) Up() (int64, int64, error) {
	return migrations.Run(svc.database, "up")
}

// UpTo runs migrations up to a certain point
func (svc *Service) UpTo(version int64) (int64, int64, error) {

	current, err := migrations.Version(svc.database)
	if err != nil {
		return -1, -1, err
	}

	if current >= version {
		return current, current, nil
	}

	vers := fmt.Sprintf("%d", version)
	return migrations.Run(svc.database, "up", vers)
}

// Down rolls back the latest migration
func (svc *Service) Down() (int64, int64, error) {
	return migrations.Run(svc.database, "down")
}

// DownTo rolls back to a specific version. Does nothing if the current version is in the future.
func (svc *Service) DownTo(version int64) (int64, int64, error) {
	current, err := migrations.Version(svc.database)
	if err != nil {
		return -1, -1, err
	}

	if current <= version {
		return current, current, nil
	}

	for i := current; i > version; i-- {
		_, new, err := migrations.Run(svc.database, "down")
		if err != nil {
			return current, new, err
		}
	}

	return current, version, nil
}
