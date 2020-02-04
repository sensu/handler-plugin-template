package main

import (
	"fmt"
	"log"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

type HandlerConfig struct {
	sensu.PluginConfig
	Example string
}

type ConfigOptions struct {
	Example sensu.PluginConfigOption
}

var (
	handlerConfig = HandlerConfig{
		PluginConfig: sensu.PluginConfig{
			Name:     "{{ .Name }}",
			Short:    "{{ .Description }}",
			Timeout:  10,
			Keyspace: "sensu.io/plugins/{{ .Name }}/config",
		},
	}

	handlerConfigOptions = ConfigOptions{
		Example: sensu.PluginConfigOption{
			Path:      "example",
			Env:       "HANDLER_EXAMPLE",
			Argument:  "example",
			Shorthand: "e",
			Default:   "",
			Usage:     "An example configuration option",
			Value:     &handlerConfig.Example,
		},
	}

	options = []*sensu.PluginConfigOption{
		&handlerConfigOptions.Example,
	}
)

func main() {
	handler := sensu.NewGoHandler(&handlerConfig.PluginConfig, options, checkArgs, executeHandler)
	handler.Execute()
}

func checkArgs(_ *types.Event) error {
	if len(handlerConfig.Example) == 0 {
		return fmt.Errorf("--example or HANDLER_EXAMPLE environment variable is required")
	}
	return nil
}

func executeHandler(event *types.Event) error {
	log.Println("executing handler with --example", handlerConfig.Example)
	return nil
}
