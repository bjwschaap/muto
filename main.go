package main

import (
	"plugin"
	"errors"

	log "github.com/sirupsen/logrus"
)

// Greeter represents the interface of our plugin
type Greeter interface {
	Greet()
}

func main() {
	// Load plugin 1
	g, err := loadPlugin("./plugin1/plugin.so")
	if err != nil {
		log.Fatal("Error loading plugin: %v", err)
	}
	// say hi with plugin 1
	g.Greet()

	// Load plugin 2
	g, err = loadPlugin("./plugin2/plugin.so")
	if err != nil {
		log.Fatal("Error loading plugin: %v", err)
	}
	// say hi with plugin 2
	g.Greet()
}


// Simple function to open/load the plugin
func loadPlugin(path string) (Greeter, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		return nil, err
	}

	greeter, ok := symGreeter.(Greeter)
	if !ok {
		return nil, errors.New("Unexpected plugin interface")
	}

	return greeter, nil
}