package amazon

import (
	"encoding/base64"
	"encoding/json"

	"github.com/Golamu/core"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	mgr "github.com/aws/aws-sdk-go/service/secretsmanager"
)

// GetAWSSecret gets a secret from AWS using the current stage
func GetAWSSecret(name, region string) (secret string, err error) {

	//Create a Secrets Manager client
	core.Debug("Creating service")
	svc := mgr.New(session.New(), aws.NewConfig().WithRegion(region))

	core.Debug("Creating input value")
	input := &mgr.GetSecretValueInput{SecretId: aws.String(name)}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/mgr/latest/apireference/API_GetSecretValue.html

	core.Debug("Pulling secret from amazon")
	result, err := svc.GetSecretValue(input)
	core.Debug("Done retrieving, processing error")
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case mgr.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				core.Error(mgr.ErrCodeDecryptionFailure, aerr.Error())
				break

			case mgr.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				core.Error(mgr.ErrCodeInternalServiceError, aerr.Error())
				break

			case mgr.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				core.Error(mgr.ErrCodeInvalidParameterException, aerr.Error())
				break

			case mgr.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				core.Error(mgr.ErrCodeInvalidRequestException, aerr.Error())
				break

			case mgr.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				core.Error(mgr.ErrCodeResourceNotFoundException, aerr.Error())
				break
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			core.Error(err.Error())
		}
		return
	}

	core.Debug("No errors")

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	if result.SecretString != nil {
		secret = *result.SecretString
		return
	}

	core.Debug("Decoding secrets")
	var size int
	byteCount := base64.StdEncoding.DecodedLen(len(result.SecretBinary))
	decodedBinarySecretBytes := make([]byte, byteCount)
	size, err = base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)

	if err != nil {
		core.Error("Base64 Decode Error:", err)
		return
	}

	core.Debug("Decoded")
	secret = string(decodedBinarySecretBytes[:size])

	return
}

// GetSecretAs takes a JSON-parseable struct pointer (!!), a secret, and a region to
// unmarshal an AWS Secret appropriately. Any errors that occur are returned unmodified
func GetSecretAs(secret string, region string, arg interface{}) error {
	confJSON, err := GetAWSSecret(secret, region)
	if err != nil {
		return err
	}

	data := []byte(confJSON)
	err = json.Unmarshal(data, &arg)
	if err != nil {
		return err
	}

	return nil
}
