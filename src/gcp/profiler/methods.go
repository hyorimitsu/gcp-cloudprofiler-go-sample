package profiler

import (
	"fmt"

	"cloud.google.com/go/profiler"
)

// Start will launch the profiler.
func (p *Profiler) Start() error {
	if p == nil {
		panic("Profiler, the receiver, is required. Please new the Profiler with profiler.New() or profiler.NewDefault().")
	}
	if err := profiler.Start(p.config); err != nil {
		return fmt.Errorf("unable to start profiler: %w", err)
	}
	return nil
}
