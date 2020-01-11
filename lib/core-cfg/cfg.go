package core_cfg

import (
	"time"
)


type DatabaseConfig struct {
	ConnectionType string `json:"connection-type" yaml:"connection-type" toml:"connection-type" xml:"connection-type"`
	User           string `json:"user-name" yaml:"user-name" toml:"user-name" xml:"user-name"`
	Password       string `json:"password" yaml:"password" toml:"password" xml:"password"`
	Host           string `json:"host" yaml:"host" toml:"host" xml:"host"`
	DatabaseName   string `json:"database-name" yaml:"database-name" toml:"database-name" xml:"database-name"`
	Charset        string `json:"charset" yaml:"charset" toml:"charset" xml:"charset"`
	ParseTime      bool   `json:"parse-time" yaml:"parse-time" toml:"parse-time" xml:"parse-time"`
	Location       string `json:"location" yaml:"location" toml:"location" xml:"location"`
	MaxIdle        int    `json:"max-idle" yaml:"max-idle" toml:"max-idle" xml:"max-idle"`
	MaxActive      int    `json:"max-active" yaml:"max-active" toml:"max-active" xml:"max-active"`
	Escaper        string `json:"escaper" yaml:"escaper" toml:"escaper" xml:"escaper"`
}

type RedisConfig struct {
	ConnectionType    string        `json:"connection-type" yaml:"connection-type" toml:"connection-type" xml:"connection-type"`
	Host              string        `json:"host" yaml:"host" toml:"host" xml:"host"`
	Password          string        `json:"password" yaml:"password" toml:"password" xml:"password"`
	Database          int           `json:"database" yaml:"database" toml:"database" xml:"database"`
	MaxIdle           int           `json:"max-idle" yaml:"max-idle" toml:"max-idle" xml:"max-idle"`
	MaxActive         int           `json:"max-active" yaml:"max-active" toml:"max-active" xml:"max-active"`
	ConnectionTimeout time.Duration `json:"connection-timeout" yaml:"connection-timeout" toml:"connection-timeout" xml:"connection-timeout"`
	WriteTimeout      time.Duration `json:"write-timeout" yaml:"write-timeout" toml:"write-timeout" xml:"write-timeout"`
	ReadTimeout       time.Duration `json:"read-timeout" yaml:"read-timeout" toml:"read-timeout" xml:"read-timeout"`
	IdleTimeout       time.Duration `json:"idle-timeout" yaml:"idle-timeout" toml:"idle-timeout" xml:"idle-timeout"`
	Wait              bool          `json:"wait" yaml:"wait" toml:"wait" xml:"wait"`
}


type DebugLogger interface {
	Debug(msg string, keyvals ...interface{})
}

func (cfg DatabaseConfig) Debug(debugLogger DebugLogger) {
	debugLogger.Debug("connected to database",
		"connection-type", cfg.ConnectionType,
		"host", cfg.Host,
		"user", cfg.User,
		"database", cfg.DatabaseName,
		"charset", cfg.Charset,
		"location", cfg.Location,
		"escaper", cfg.Escaper,
	)

}


