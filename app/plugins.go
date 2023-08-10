package app

import (
	"github.com/sonrhq/core/internal/highway"
)

// EnablePlugins enables the plugins.
func EnablePlugins() {
	highway.StartAPI()
	highway.StartDB()
	highway.StartSQL()
}
