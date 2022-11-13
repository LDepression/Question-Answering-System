package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	v1 "wenba/internal/api/v1"
	"wenba/internal/global"
	"wenba/internal/route/router"
	"wenba/internal/settings"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}
func initSettings() {
	settings.AllInit()
}
func main() {
	initSettings()
	if global.Settings.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := router.NewRouter() //注册路由
	if err := v1.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans ,err:%v\n", err)
		return
	}
	s := http.Server{
		Addr:           global.Settings.Server.Addr,
		Handler:        r,
		ReadTimeout:    global.Settings.Server.ReadTimeOut,
		WriteTimeout:   global.Settings.Server.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Info("server start suc")
	fmt.Println("AppName:", global.Settings.App.Name, "Version:", global.Settings.App.Version, "Addr:", global.Settings.Server.Addr)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			global.Logger.Info(err.Error())
		}
	}()
	gracefulExit(&s) // 优雅退出
	global.Logger.Info("Server exited!")
}

// 优雅退出
func gracefulExit(s *http.Server) {
	// 退出通知
	quit := make(chan os.Signal, 1)
	// 等待退出通知
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("ShutDown Server...")
	// 给几秒完成剩余任务
	ctx, cancel := context.WithTimeout(context.Background(), global.Settings.Server.DefaultContextTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil { // 优雅退出
		global.Logger.Info("Server forced to ShutDown,Err:" + err.Error())
	}
}
