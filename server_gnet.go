package gmqtt

import (
	"fmt"
	"sync/atomic"

	"github.com/lei006/gmqtt/protocol"

	"github.com/panjf2000/gnet/v2"
	"github.com/sohaha/zlsgo/zlog"
)

type HandleCallback func(con gnet.Conn) error

type gnetServer struct {
	gnet.BuiltinEventEngine
	eng          gnet.Engine
	Network      string
	Addr         string
	Multicore    bool
	connected    int32
	disconnected int32
	clientMsr    clientManager

	callback HandleCallback
}

func MakeGnetServer() *gnetServer {
	return &gnetServer{}
}

func (s *gnetServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	zlog.Infof("mqtt running server on %s with multi-core=%t\n", fmt.Sprintf("%s://%s", s.Network, s.Addr), s.Multicore)
	s.eng = eng
	return
}

func (srv *gnetServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(new(protocol.SimpleCodec))
	zlog.Success("xxxxxxxxxxxxxxxx", c.LocalAddr().String(), c.RemoteAddr().String())
	atomic.AddInt32(&srv.connected, 1)
	out = []byte("sweetness\r\n")

	srv.clientMsr.Add()

	return
}

func (s *gnetServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		zlog.Infof("error occurred on connection=%s, %v\n", c.RemoteAddr().String(), err)
	}
	disconnected := atomic.AddInt32(&s.disconnected, 1)
	connected := atomic.AddInt32(&s.connected, -1)
	if connected == 0 {
		zlog.Infof("all %d connections are closed, shut it down\n", disconnected)
		action = gnet.Close
	}
	return
}

func (srv *gnetServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	if srv.callback != nil {
		srv.callback(c)
	}
	action = gnet.Close

	/*
		codec := c.Context().(*protocol.SimpleCodec)
		var packets [][]byte
		for {
			data, err := codec.Decode(c)
			if err == protocol.ErrIncompletePacket {
				break
			}

			if err != nil {
				zlog.Errorf("invalid packet: %v\n", err)
				return gnet.Close
			}
			packet, _ := codec.Encode(data)
			packets = append(packets, packet)
		}
		if n := len(packets); n > 1 {
			_, _ = c.Writev(packets)
		} else if n == 1 {
			_, _ = c.Write(packets[0])
		}
	*/
	return
}

func (gnet_srv *gnetServer) ListenAndServer(addr string) error {

	err := gnet.Run(gnet_srv, "tcp://"+addr, gnet.WithMulticore(true))
	zlog.Infof("server exits with error: %v", err)

	return err
}

func (gnet_srv *gnetServer) OnData(callback HandleCallback) error {

	gnet_srv.callback = callback

	return nil
}
