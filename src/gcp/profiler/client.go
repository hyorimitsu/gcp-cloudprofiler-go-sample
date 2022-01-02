package profiler

import (
	"os"

	"cloud.google.com/go/profiler"
	"google.golang.org/api/option"
)

// Profiler is Cloud Profiler wrapper.
type Profiler struct {
	config  profiler.Config
	options []option.ClientOption
}

// New returns new Profiler.
func New(config profiler.Config, options ...option.ClientOption) Profiler {
	return Profiler{
		config,
		options,
	}
}

// NewDefault returns new default Profiler.
func NewDefault() Profiler {
	return Profiler{
		config: profiler.Config{
			Service:        serviceName(),
			ServiceVersion: serviceVersion(),
			// see below for other settings.
			//   https://github.com/googleapis/google-cloud-go/blob/main/profiler/profiler.go#L106-L188
		},
	}
}

func serviceName() string {
	if val := os.Getenv("SERVICE_NAME"); val != "" {
		return val
	}
	return DefaultServiceName
}

func serviceVersion() string {
	if val := os.Getenv("SERVICE_VERSION"); val != "" {
		return val
	}
	return DefaultServiceVersion
}
