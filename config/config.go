package config

import "os"

// Port returns the port number to open.
func Port() string {
	return os.Getenv("PORT")
}
