package pg

import (
	"encoding/json"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/Golamu/core/amazon"
)

// DB is a re-export of the go-pg DB
type DB = pg.DB

// Conn is a re-export of the go-pg Conn
type Conn = pg.Conn

// ConnectWithSecret first calls GetConnectionSecret and returns the error if there is one
// If it is successful, it calls connect, and returns the connection to your for your use
func ConnectWithSecret(region string, secret string) (db *DB, err error) {
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
	creds = DatabaseCredentials{Region: region}
	secret, err := amazon.GetAWSSecret(secretName, region)

	if err != nil {
		fmt.Printf("Unable to get database credentials because:\n%s", err.Error())
		return
	}

	err = json.Unmarshal([]byte(secret), &creds)
	if err != nil {
		fmt.Printf("An error occurred unmarshalling DB credentials: \n%s", err.Error())
		return
	}

	return
}
