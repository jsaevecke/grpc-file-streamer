package logging

import (
	"errors"

	"grpc-file-streamer/cerrors"

	"github.com/asaskevich/govalidator"
	"github.com/rs/zerolog"
)

func EnrichLoggerWithServiceFields(serviceFields *ParamsServiceFields, logger *zerolog.Logger) (zerolog.Logger, error) {
	if serviceFields == nil {
		return *logger, cerrors.ErrValidation{
			Issue:  errors.New("service fields are nil"),
			Caller: "EnrichLoggerWithServiceFields",
		}
	}

	if _, errValidate := govalidator.ValidateStruct(serviceFields); errValidate != nil {
		return *logger, cerrors.ErrValidation{
			Issue:  errValidate,
			Caller: "EnrichLoggerWithServiceFields",
		}
	}

	return logger.
		With().
		Str(LogFieldServiceGroup.String(), serviceFields.ServiceGroup).
		Str(LogFieldServiceName.String(), serviceFields.ServiceName).
		Str(LogFieldServiceType.String(), serviceFields.ServiceType).
		Str(LogFieldServiceVersion.String(), serviceFields.ServiceVersion).
		Logger(), nil
}
