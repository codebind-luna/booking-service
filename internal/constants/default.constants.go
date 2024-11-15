package constants

const (
	DefaultGRPCHost                   string = "localhost"
	DefaultGRPCPort                   int    = 50051
	MinPort                           int    = 1024
	MaxPort                           int    = 65535
	DefaultShutdownContextWaitSeconds        = 5
)

const (
	DefaultLogLevel  string = "info"
	DefaultLogFormat string = "text"
)

const (
	DefaultRepositoryType string = "in-memory"
)

const (
	DefaultSeats int = 10
)
