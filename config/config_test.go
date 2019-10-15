package config

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	// "github.com/pelletier/go-toml"
	btoml "github.com/BurntSushi/toml"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

func TestJSON(t *testing.T) {
	var config ServerConfig

	f, err := os.Open("server-config.example.json")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
		return
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(config)
}

func TestYAML(t *testing.T) {
	var config ServerConfig

	f, err := os.Open("server-config.example.yml")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
		return
	}
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(config)
}

func TestTOML(t *testing.T) {
	var config ServerConfig

	f, err := os.Open("server-config.example.toml")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
		return
	}
	err = toml.Unmarshal(b, &config)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(config)
}

func TestBTOML(t *testing.T) {
	var config ServerConfig

	f, err := os.Open("server-config.example.toml")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = btoml.Decode(string(b), &config)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(config)
}

func TestXML(t *testing.T) {
	var config ServerConfig

	f, err := os.Open("server-config.example.xml")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
		return
	}
	err = xml.Unmarshal(b, &config)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(config)
}
