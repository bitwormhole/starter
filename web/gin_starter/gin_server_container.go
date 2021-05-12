package gin_starter

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter/application"
	"github.com/gin-gonic/gin"
)

type GinWebController interface {
	Config(engine *gin.Engine) error
}

type GinServerContainer struct {
	port           int
	host           string
	engine         *gin.Engine
	controllers    []GinWebController
	runtimeContext application.RuntimeContext
}

func (inst *GinServerContainer) Inject(context application.RuntimeContext) error {

	components := context.GetComponents()
	keys := components.GetComponentNameList(false)
	ctrlist := []GinWebController{}

	for index := range keys {
		key := keys[index]
		com, err := components.GetComponent(key)
		if err != nil {
			return err
		}
		controller, ok := com.(GinWebController)
		if ok {
			ctrlist = append(ctrlist, controller)
		}
	}

	inst.loadProperties(context)
	inst.controllers = ctrlist
	return nil
}

func (inst *GinServerContainer) loadProperties(context application.RuntimeContext) error {

	getter := context.GetProperties()
	str_host := getter.GetProperty("server.host", "")
	str_port := getter.GetProperty("server.port", "8080")

	port, err := strconv.ParseInt(str_port, 10, 32)
	if err != nil {
		return err
	}
	if port < 1 || 65535 < port {
		return errors.New("bad port value .")
	}

	inst.port = int(port)
	inst.host = str_host
	return nil
}

func (inst *GinServerContainer) initEngine() error {

	server := gin.Default()

	server.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello":  "world",
			"server": "gin",
			"port":   666,
		})
	})

	inst.engine = server
	return nil
}

func (inst *GinServerContainer) initControllers() error {

	controllers := inst.controllers
	engine := inst.engine

	for index := range controllers {
		ctrl := controllers[index]
		err := ctrl.Config(engine)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *GinServerContainer) initStatic() error {

	res := inst.runtimeContext.GetResources()
	engine := inst.engine
	adapter := &GinFsAdapter{res: res}

	engine.StaticFS("/", http.FS(adapter.GetFS()))

	return nil
}

func (inst *GinServerContainer) Run() error {

	port := inst.port
	host := inst.host

	if port < 1 {
		port = 8080
	}

	addr := host + ":" + strconv.Itoa(port)
	inst.engine.Run(addr)
	return nil
}

func (inst *GinServerContainer) Init() error {

	err := inst.initEngine()
	if err != nil {
		return err
	}

	err = inst.initControllers()
	if err != nil {
		return err
	}

	return nil
}

func (inst *GinServerContainer) Destroy() error {

	return nil
}
