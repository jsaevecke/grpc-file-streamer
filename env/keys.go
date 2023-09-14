package env

// keys to retrieve environmental variables
const (
	// environment
	EnvEnvironment = "ENVIRONMENT"

	// service
	EnvServiceName    = "SERVICE_NAME"
	EnvServiceType    = "SERVICE_TYPE"
	EnvServiceGroup   = "SERVICE_GROUP"
	EnvServiceVersion = "SERVICE_VERSION"

	// logging
	EnvPrettyPrint = "PRETTY_PRINT"
	EnvLogLevel    = "LOG_LEVEL"
)

const (
	ValEnvironmentLocal = "local"
)
