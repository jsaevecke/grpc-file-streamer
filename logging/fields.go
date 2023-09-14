package logging

type LogField string

func (c LogField) String() string {
	return string(c)
}

// Logger fields for setting or retrieving contextual information
const (
	// Environment
	LogFieldEnvironment LogField = "environment"

	// Service
	LogFieldServiceName    LogField = "service"
	LogFieldServiceType    LogField = "service_type"
	LogFieldServiceGroup   LogField = "service_group"
	LogFieldServiceVersion LogField = "service_version"

	// Time
	LogFieldStartTime LogField = "start_time"
	LogFieldEndTime   LogField = "end_time"
)
