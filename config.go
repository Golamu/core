package core

// IAppConfigService is a service that allows any other service to pull configuration
// variables without directly relying on the process.env, and can instead pull from
// things like the AWS Secrets manager
type IAppConfigService interface {
	Get(key string) string
	Set(key string, value string)
}
