package gcp

import "os"

// RunningOnGCP returns true if the application is running on GCP.
func RunningOnGCP() bool {
	return os.Getenv("RUNNING_ON_GCP") != ""
}
