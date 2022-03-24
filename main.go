package main

import (
	"fmt"
	"gin-base/api/v1/model"
	"gin-base/router"
	"gin-base/tool"
	"net/http"
	"time"
)

func main() {

	// router := routers.InitRouter()
	s := &http.Server{

		Addr: fmt.Sprintf(":%d", tool.GetConfig().Server.ServerPort),
		// Handler:        router,
		ReadTimeout:    time.Duration(tool.GetConfig().Server.ServerReadTimeout * int(time.Second)),
		WriteTimeout:   time.Duration(tool.GetConfig().Server.ServerWriteTimeout * int(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func init() {

	tool.ParseConfig()
	router.InitRouter()
	model.NewDbEngine()

}
