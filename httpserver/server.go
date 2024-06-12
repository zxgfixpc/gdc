package httpserver

import (
	"context"
	"errors"
	"fmt"
	"gdc/conf"
	"gdc/lib/infra"
	"gdc/lib/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Start() {
	ctx := context.TODO()

	config := &conf.ServerConf{}
	if err := conf.Parser(config, conf.ServerPath); err != nil {
		panic(err)
	}

	g := gin.New()

	g.Use(log.LoggerMiddleware())
	g.Use(gin.Recovery())

	// 初始化基础组件
	if err := infra.Start(ctx); err != nil {
		panic(err)
	}

	registerRouter(g)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: g,
	}
	go func() {
		fmt.Println("listen port:", config.Port)
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Panic(ctx, err)
			}
		}
	}()

	quitSig := make(chan os.Signal, 1)
	signal.Notify(quitSig, syscall.SIGINT, syscall.SIGTERM) //等待退出信号
	<-quitSig

	if err := server.Shutdown(ctx); err != nil {
		log.ErrorF(ctx, "shutdown err:%v", err)
	} else {
		log.Info(ctx, "shutdown success")
	}

	quitSvr := make(chan bool, 1)
	go shutdown(ctx, quitSvr) //优雅关闭，比如一些长连接等

	select {
	case <-time.After(time.Duration(config.ShutDownWait) * time.Second):
		log.Error(ctx, "shutdown timeout")
	case <-quitSvr:
		log.Info(ctx, "shutdown success")
	}
}

func shutdown(ctx context.Context, c chan bool) {
	defer func() {
		c <- true
	}()

	// 自定义退出

	// 关闭基础组件等长连接
	infra.Shutdown(ctx)
}
