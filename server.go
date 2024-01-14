package gmqtt

import (
	"sync"
	"time"

	"github.com/sohaha/zlsgo/zlog"
)

// 创建服务器
func ListenAndServer(config *ServerConfig, wg *sync.WaitGroup) {

	defer wg.Done()

	for {

		srv := MakeMqttServer(config)
		err := srv.ListenAndServer()
		zlog.Error("err:", err)

		time.Sleep(300 * time.Second)
	}
}
