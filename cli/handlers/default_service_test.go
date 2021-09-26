package handlers

import (
	"testing"

	"github.com/bitwormhole/starter/cli"
	"github.com/bitwormhole/starter/cli/filters"
)

func TestDefaultService(t *testing.T) {

	service := &cli.DefaultSerivce{}

	service.AddFilter(50, &filters.ContextFilter{Service: service})
	service.AddFilter(45, &filters.MultilineSupportFilter{})
	service.AddFilter(40, &filters.HandlerFinderFilter{})
	service.AddFilter(30, &filters.ExecutorFilter{})
	service.AddFilter(0, &filters.NopFilter{})

	service.RegisterHandler("help", &Help{})

	client := service.GetClient()

	err := client.ExecuteScript("help hello world")
	if err != nil {
		t.Error(err)
	}
}
