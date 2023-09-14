package logging

// ParamsServiceFields contains the minimum required information about the service using the logger.
type ParamsServiceFields struct {
	ServiceGroup   string `valid:"required"`
	ServiceName    string `valid:"required"`
	ServiceType    string `valid:"required"`
	ServiceVersion string `valid:"required"`
}
