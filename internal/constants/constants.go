package constants

const (

	// Server constants
	ServerPort = 8080

	// Logger constants
	RequestID       = "request_id"
	EnvLogLevel     = "LOG_LEVEL"
	EnvLogStdOut    = "LOG_STDOUT"
	EnvLogFormatter = "LOG_FORMATTER"

	// Path constants
	PingBasePath = "ping"
	MockBasePath = "mock"

	// Error constants
	BadRequestMessage          = "bad request"
	NotFoundMessage            = "not found"
	InternalServerErrorMessage = "internal server error"
)
