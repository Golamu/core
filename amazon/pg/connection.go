package pg

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/Golamu/core"
	"github.com/Golamu/core/amazon"
	"github.com/go-pg/pg/v10"
)

// DB is a re-export of the go-pg DB
type DB = pg.DB

// Conn is a re-export of the go-pg Conn
type Conn = pg.Conn

// ConnectWithSecret first calls GetConnectionSecret and returns the error if there is one
// If it is successful, it calls connect, and returns the connection to your for your use
func ConnectWithSecret(region string, secret string) (db *DB, err error) {
	core.Log("Getting connection secret %s, region %s", secret, region)
	creds, err := GetConnectionSecret(region, secret)
	if err != nil {
		return
	}

	return creds.Connect()
}

// GetConnectionSecret automatically populates a new DatabaseCredentials struct based on the
// region and secret arguments
func GetConnectionSecret(
	region string,
	secretName string,
) (creds DatabaseCredentials, err error) {
	core.Debug("Retrieving secret")
	secret, err := amazon.GetAWSSecret(secretName, region)

	core.Debug("Secret retrieved:\n\t %s", secret)

	if err != nil {
		core.Error("Unable to get database credentials because:\n%s", err.Error())
		return
	}

	creds = DatabaseCredentials{Region: region}
	err = json.Unmarshal([]byte(secret), &creds)
	if err != nil {
		core.Error("An error occurred unmarshalling DB credentials: \n%s", err.Error())
		return
	}

	creds.UseIAM = os.Getenv("DB_USE_IAM") == "true"

	msgs := []string{
		"Retrieved:",
		"  Database:   %s",
		"  Endpoint:   %s",
		"  Username:   %s",
		"  UseIAM:     %t",
		"  Region:     %s",
		"  Port:       %d",
		"  AppName:    %s",
	}

	core.Debug(
		strings.Join(msgs, "\n"),
		creds.Database,
		creds.Host,
		creds.UserName,
		creds.UseIAM,
		creds.Region,
		creds.Port,
		creds.ApplicationName,
	)

	return
}
