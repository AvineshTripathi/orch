package models

// Config represents common configuration fields for all plugins
type Config struct {
	Cfg any
	LogFile  string
	LogLevel int
}

// TaskPlugin is the interface that plugins must implement
type TaskPlugin interface {
	Run(*Config) error
}