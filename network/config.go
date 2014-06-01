package network

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Port        int            `json:"port"`
	TemplateDir string         `json:"templates"`
	StaticDir   string         `json:"static"`
	Database    DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Return a string of credentials approriate for Go's sql.Open() func
func (db DatabaseConfig) Credentials() string {
	// TODO Different credentials for different drivers
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		db.Host,
		db.Port,
		db.Name,
		db.User,
		db.Password,
	)
}

func Parse() (config Config) {
	// It's okay if this fails
	if f, err := os.Open("./settings.json"); err == nil {
		if b, err := ioutil.ReadAll(f); err == nil {
			if err = json.Unmarshal(b, &config); err != nil {
				panic(err)
			}
		}
	}
	flag.IntVar(&config.Port, "port", 8001, "port for the HTTP server")
	flag.StringVar(
		&config.TemplateDir,
		"templates",
		"./templates",
		"directory for templates",
	)
	flag.StringVar(
		&config.StaticDir,
		"static",
		"./static",
		"directory for static files",
	)
	flag.Parse()
	return
}
