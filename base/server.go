package base

import (
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/mszhangyi/infra"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

var Service common.Service

func Server() common.Service {
	Check(Service)
	return Service
}

type ServerStarter struct {
	infra.BaseStarter
}

func (i *ServerStarter) Init() {
	Service = daprd.NewService(":" + strconv.Itoa(props.Port))
}

func (i *ServerStarter) Start() {
	go func() {
		// 服务连接
		if err := Service.Start(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("error listenning: %v", err)
		}
	}()
	logrus.Debug("服务器正在运行,端口：" + strconv.Itoa(props.Port))
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Print("Shutdown Server ...")
}
