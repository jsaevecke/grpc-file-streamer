package cerrors

const (
	// Initialization Error Messages
	ErrorMsgFailedLoggerInitialzation = "failed to initialize logger"
	ErrorMsgFailedLoggerEnrichment    = "failed to enrich logger with service fields"
	ErrorMsgFailedInitConfiguration   = "failed to initialize configuration"

	// Service Error Messages
	ErrorMsgFailedCreateService = "failed to create service: %s"
	ErrorMsgFailedStartService  = "failed to start service: %s"
)

const _space = " "

var (
	OSExitForConfigurationIssues  = 1
	OSExitForServiceIssues        = 2
	OSExitForApplicationIssues    = 3
	OSExitForFileOperationsIssues = 4
	OSExitForGRPCIssues           = 5
	OSExitForLoggingIssues        = 6
)
