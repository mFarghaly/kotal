package v1alpha1

const (
	// DefaultClient is the default ethereum 2.0 client
	DefaultClient = TekuClient
	// DefaultRestPort is the default Beacon REST api port
	DefaultRestPort uint = 5051
	// DefaultRPCPort is the default RPC server port
	DefaultRPCPort uint = 4000
	// DefaultGRPCPort is the default GRPC gateway server port
	DefaultGRPCPort uint = 3500
	// DefaultRPCHost is the default host on which RPC server should listen
	DefaultRPCHost = "0.0.0.0"
)
