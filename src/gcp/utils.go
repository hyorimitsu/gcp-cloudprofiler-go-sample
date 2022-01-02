package gcp

import "os"

// RunningOnGCP returns true if the application is running on GCP.
func RunningOnGCP() bool {
	// this environment variable is automatically set
	// when you run application on GCP
	return os.Getenv("GOOGLE_CLOUD_PROJECT") != ""
}
