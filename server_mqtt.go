package gmqtt

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/sohaha/zlsgo/zlog"
)

type MqttServer struct {
	config *ServerConfig
	netSrv *gnetServer
}

func MakeMqttServer(config *ServerConfig) *MqttServer {
	srv := &MqttServer{
		config: config,
	}

	srv.netSrv = MakeGnetServer()

	return srv
}

func (mqtt_srv *MqttServer) ListenAndServer() error {

	mqtt_srv.netSrv.OnData(mqtt_srv.handleCallback)

	return mqtt_srv.netSrv.ListenAndServer(mqtt_srv.config.Addr)
}
func (mqtt_srv *MqttServer) handleCallback(con gnet.Conn) error {

	zlog.Debug("handleCallback   ")

	return nil
}
