package pg

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	"github.com/go-pg/pg/v10"
)

// DatabaseCredentials matches the type found in the secrets manager for Postgres secrets.
// This exposes serialization and AWS -> pg.Options functionality
type DatabaseCredentials struct {
	UserName        string `json:"username"`
	Engine          string `json:"engine"`
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Database        string `json:"dbname"`
	Password        string `json:"password"`
	ID              string `json:"dbInstanceIdentifier"`
	Region          string
	ApplicationName string `json:"appname"`
	UseIAM          bool
}

// Endpoint returns the host:port for this credentials object
func (cr *DatabaseCredentials) Endpoint() string {
	return fmt.Sprintf("%s:%d", cr.Host, cr.Port)
}

// IAMPassword generates the AWS password using IAM
func (cr *DatabaseCredentials) IAMPassword() (string, error) {
	awsCreds := credentials.NewEnvCredentials()
	authToken, err := rdsutils.BuildAuthToken(cr.Endpoint(), cr.Region, cr.UserName, awsCreds)
	if err != nil {
		return "", err
	}

	return authToken, nil
}

// Connect takes these credentials and creates a Postgres connection out of them
func (cr *DatabaseCredentials) Connect() (db *DB, err error) {
	pwd := cr.Password
	if cr.UseIAM {
		pwd, err = cr.IAMPassword()
	}

	if err != nil {
		return nil, err
	}

	opts := pg.Options{
		User:            cr.UserName,
		Addr:            cr.Endpoint(),
		Database:        cr.Database,
		Password:        pwd,
		ApplicationName: cr.ApplicationName,
		OnConnect:       onSuccessfulConnect,
		TLSConfig:       &tls.Config{InsecureSkipVerify: true},
	}

	db = pg.Connect(&opts)

	return
}

func onSuccessfulConnect(ctx context.Context, conn *Conn) error {
	fmt.Println("Database: Successfully connected")
	return nil
}
