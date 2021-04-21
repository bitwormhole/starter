package gin_starter

import (
	"github.com/bitwormhole/starter/application"
)

func Run(context application.RuntimeContext) error {
	pServer, err := context.GetComponents().GetComponent(ID_GIN_WEB_SERVER)
	if err != nil {
		return err
	}
	container := pServer.(*GinServerContainer)
	return container.Run()
}
