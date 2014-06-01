package antisocial

import (
	"flag"
)

type Config struct {
	Port        int
	TemplateDir string
	StaticDir   string
}

func Parse() (config Config) {
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
