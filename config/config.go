package config

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"time"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

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

type DatabaseConfig struct {
	ConnectionType string `json:"connection-type" yaml:"connection-type" toml:"connection-type" xml:"connection-type"`
	User           string `json:"user-name" yaml:"user-name" toml:"user-name" xml:"user-name"`
	Password       string `json:"password" yaml:"password" toml:"password" xml:"password"`
	Host           string `json:"host" yaml:"host" toml:"host" xml:"host"`
	DatabaseName   string `json:"database-name" yaml:"database-name" toml:"database-name" xml:"database-name"`
	Charset        string `json:"charset" yaml:"charset" toml:"charset" xml:"charset"`
	ParseTime      bool   `json:"parse-time" yaml:"parse-time" toml:"parse-time" xml:"parse-time"`
	Location       string `json:"location" yaml:"location" toml:"location" xml:"location"`
}

type Label struct {
	Key   string `json:"key" yaml:"key" toml:"key" xml:"key"`
	Value string `json:"value" yaml:"value" toml:"value" xml:"value"`
}

type ServerConfig struct {
	LoadType       string         `json:"-" yaml:"-" toml:"-" xml:"-"`
	Name           xml.Name       `json:"-" yaml:"-" toml:"-" xml:"server-config"`
	Labels         []Label        `json:"label" yaml:"label" toml:"label" xml:"label"`
	DatabaseConfig *DatabaseConfig `json:"database" yaml:"database" toml:"database" xml:"database"`
	RedisConfig    *RedisConfig    `json:"redis" yaml:"redis" toml:"redis" xml:"redis"`
}

func Load(config *ServerConfig, configpath string) error {
	for _, configX := range []struct {
		Type      string
		Unmarshal func([]byte, interface{}) error
	}{
		{".json", json.Unmarshal}, {".yml", yaml.Unmarshal},
		{".toml", toml.Unmarshal}, {".xml", xml.Unmarshal}} {
		if _, err := os.Stat(configpath + configX.Type); err == nil {
			f, err := os.Open(configpath + configX.Type)
			if err != nil {
				return err
			}

			b, err := ioutil.ReadAll(f)
			_ = f.Close()
			if err != nil {
				return err
			}
			err = configX.Unmarshal(b, config)
			if err != nil {
				return err
			}
			config.LoadType = configX.Type
			return nil
		}
	}

	return errors.New("no such file in the root directory")
}

func LoadStatic(config interface{}, configpath string) error {
	for _, configX := range []struct {
		Type      string
		Unmarshal func([]byte, interface{}) error
	}{
		{".json", json.Unmarshal}, {".yml", yaml.Unmarshal},
		{".toml", toml.Unmarshal}, {".xml", xml.Unmarshal}} {
		if _, err := os.Stat(configpath + configX.Type); err == nil {
			f, err := os.Open(configpath + configX.Type)
			if err != nil {
				return err
			}

			b, err := ioutil.ReadAll(f)
			_ = f.Close()
			if err != nil {
				return err
			}
			err = configX.Unmarshal(b, config)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("no such file in the root directory")
}

func Save(config *ServerConfig, configpath string) error {
	var b []byte
	var err error
	switch config.LoadType {
	case ".json":
		b, err = json.MarshalIndent(config, "", "    ")
		if err != nil {
			return err
		}
	case ".yml":
		b, err = yaml.Marshal(config)
		if err != nil {
			return err
		}
	case ".toml":
		b, err = toml.Marshal(config)
		if err != nil {
			return err
		}
	case ".xml":
		b, err = xml.MarshalIndent(config, "", "    ")
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(configpath + config.LoadType); err == nil {
		f, err := os.OpenFile(configpath+config.LoadType, os.O_WRONLY|os.O_TRUNC, 0333)
		if err != nil {
			return err
		}

		_, err = f.Write(b)
		_ = f.Close()
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("no such file in the root directory")
}
