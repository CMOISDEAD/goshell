package internals

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

type Alias struct {
	Alias   string
	Command string
}

type Configuration struct {
	Char    interface{}
	Scripts []string
}

func Config(dev bool) Configuration {
	path := "/home/camilo/goshell.toml"
	if dev {
		path = "./goshell.toml"
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	config, _ := toml.Load(string(data))
	var scripts []string
	for _, value := range config.Get("init.scripts").([]interface{}) {
		scripts = append(scripts, value.(string))
	}
	return Configuration{
		Char:    config.Get("theme.prompt"),
		Scripts: scripts,
	}
}

func GetAlias(dev bool) []Alias {
	var alias []Alias
	path := "/home/camilo/goshell.toml"
	if dev {
		path = "./goshell.toml"
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	config, _ := toml.Load(string(data))
	for _, value := range config.Get("alias").([]*toml.Tree) {
		aux := Alias{
			Alias:   value.Get("alias").(string),
			Command: value.Get("command").(string),
		}
		alias = append(alias, aux)
	}
	return alias
}
