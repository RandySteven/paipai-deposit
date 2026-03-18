// Package config provides configuration management for the application.
// It handles loading, parsing, and accessing configuration values from YAML files
// and supports server, database (PostgreSQL, MongoDB), Redis, NSQ, and Temporal settings.
package configs

import "time"

type (
	// Config represents the root configuration structure for the application.
	// It contains nested configuration for all services including server, databases,
	// message queues, and workflow orchestration.
	Config struct {
		Configs struct {
			// Server contains HTTP server configuration including host, port, and timeouts.
			Server struct {
				Host    string `yaml:"host"`
				Port    string `yaml:"port"`
				Timeout struct {
					Server time.Duration `yaml:"server"`
					Read   time.Duration `yaml:"read"`
					Write  time.Duration `yaml:"write"`
					Idle   time.Duration `yaml:"idle"`
				} `yaml:"timeout"`
			} `yaml:"server"`

			// Postgres contains PostgreSQL database connection settings.
			Postgres struct {
				Host   string `yaml:"host"`
				Port   string `yaml:"port"`
				DbName string `yaml:"dbname"`
				DbUser string `yaml:"dbuser"`
				DbPass string `yaml:"dbpass"`
			} `yaml:"postgres"`

			// Mongodb contains MongoDB connection settings.
			Mongodb struct {
				User     string `yaml:"user"`
				Password string `yaml:"password"`
				Host     string `yaml:"host"`
				DbName   string `yaml:"dbname"`
			} `yaml:"mongodb"`

			// Redis contains Redis cache connection settings.
			Redis struct {
				Host     string `yaml:"host"`
				Port     string `yaml:"port"`
				Password string `yaml:"password"`
			} `yaml:"redis"`

			// NSQ contains NSQ message queue settings.
			NSQ struct {
				Host            string        `yaml:"host"`
				Port            string        `yaml:"port"`
				Topic           string        `yaml:"topic"`
				Channel         string        `yaml:"channel"`
				MaxInFlight     int           `yaml:"maxInFlight"`
				MaxRequeueDelay time.Duration `yaml:"maxRequeueDelay"`
				MaxRequeueCount int           `yaml:"maxRequeueCount"`
			} `yaml:"nsq"`

			// Temporal contains Temporal workflow orchestration settings.
			Temporal struct {
				Host          string `yaml:"host"`
				Port          string `yaml:"port"`
				TaskQueue     string `yaml:"task_queue"`
				Namespace     string `yaml:"namespace"`
				WorkerOptions *struct {
					MaxConcurrentActivityExecutionSize      int     `yaml:"maxConcurrentActivityExecutionSize"`
					WorkerActivitiesPerSecond               float64 `yaml:"workerActivitiesPerSecond"`
					MaxConcurrentLocalActivityExecutionSize int     `yaml:"maxConcurrentLocalActivityExecutionSize"`
					WorkerLocalActivitiesPerSecond          float64 `yaml:"workerLocalActivitiesPerSecond"`
				} `yaml:"workerOptions"`
			} `yaml:"temporal"`
		} `yaml:"configs"`
	}
)

// GetConfigs returns the nested Configs struct containing all service configurations.
func (c *Config) GetConfigs() any {
	return &c.Configs
}
