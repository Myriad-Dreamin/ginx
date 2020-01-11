package core_cfg


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


