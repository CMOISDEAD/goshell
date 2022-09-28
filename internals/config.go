package internals

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

type Configuration struct {
	Char  interface{}
	Scripts []string
}

func Config() Configuration {
	data, err := os.ReadFile("./goshell.toml")
	if err != nil {
		fmt.Println(err)
	}
	config, _ := toml.Load(string(data))
	var scripts []string
	for _, value := range config.Get("init.scripts").([] interface {}) {
		scripts = append(scripts, value.(string))
	}
	return Configuration {
		Char : config.Get("theme.prompt"),
		Scripts: scripts,
	}
}
