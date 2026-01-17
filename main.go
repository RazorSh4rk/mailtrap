package main

import (
	"fmt"

	"github.com/phires/go-guerrilla"
	"github.com/phires/go-guerrilla/backends"
	"github.com/phires/go-guerrilla/mail"
)

var LoggerProcessor = func() backends.Decorator {
	initFunc := backends.InitializeWith(func(backendConfig backends.BackendConfig) error {
		backends.Log().Info("Logger processor loaded")
		return nil
	})
	backends.Svc.AddInitializer(initFunc)
	return func(p backends.Processor) backends.Processor {
		return backends.ProcessWith(
			func(e *mail.Envelope, task backends.SelectTask) (backends.Result, error) {
				backends.Log().Info(e.String())
				return p.Process(e, task)
			},
		)
	}
}

func main() {
	cfg := &guerrilla.AppConfig{
		Servers: []guerrilla.ServerConfig{
			{
				ListenInterface: "0.0.0.0:25",
				IsEnabled:       true,
			},
		},
		AllowedHosts: []string{
			"*",
		},
		BackendConfig: backends.BackendConfig{
			"save_process":     "HeadersParser|Debugger|FullLogger",
			"validate_process": "FullLogger",
		},
	}

	d := guerrilla.Daemon{Config: cfg}
	d.AddProcessor("FullLogger", LoggerProcessor)
	err := d.Start()

	if err == nil {
		fmt.Println("Server Started!")
	}

	for {
	}
}
