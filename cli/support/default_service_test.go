package support

import (
	"testing"

	"github.com/bitwormhole/starter/cli/filters"
	"github.com/bitwormhole/starter/cli/handlers"
)

func TestDefaultService(t *testing.T) {

	service := &DefaultSerivce{}

	service.AddFilter(50, &filters.ContextFilter{Service: service})
	service.AddFilter(45, &filters.MultilineSupportFilter{})
	service.AddFilter(40, &filters.HandlerFinderFilter{})
	service.AddFilter(30, &filters.ExecutorFilter{})
	service.AddFilter(0, &filters.NopFilter{})

	service.RegisterHandler("help", &handlers.Help{})

	// client := service.GetClient(nil)
	// err := client.ExecuteScript("help hello world")
	// if err != nil {
	// 	t.Error(err)
	// }
}
