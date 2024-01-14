package gmqtt

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/lei006/gmqtt/protocol"

	"github.com/panjf2000/gnet/v2"
	"github.com/sohaha/zlsgo/zlog"
)

type EchoServer struct {
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	Addr      string
	Multicore bool
}

func (es *EchoServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	log.Printf("echo server with multi-core=%t is listening on %s\n", es.Multicore, es.Addr)
	return gnet.None
}

func (es *EchoServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	c.Write(buf)
	return gnet.None
}

type SimpleServer struct {
	gnet.BuiltinEventEngine
	eng          gnet.Engine
	Network      string
	Addr         string
	Multicore    bool
	connected    int32
	disconnected int32
}

func (s *SimpleServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	zlog.Infof("running server on %s with multi-core=%t",
		fmt.Sprintf("%s://%s", s.Network, s.Addr), s.Multicore)
	s.eng = eng
	return
}

func (s *SimpleServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(new(protocol.SimpleCodec))
	atomic.AddInt32(&s.connected, 1)
	out = []byte("sweetness\r\n")
	return
}

func (s *SimpleServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		zlog.Infof("error occurred on connection=%s, %v\n", c.RemoteAddr().String(), err)
	}
	disconnected := atomic.AddInt32(&s.disconnected, 1)
	connected := atomic.AddInt32(&s.connected, -1)
	if connected == 0 {
		zlog.Infof("all %d connections are closed, shut it down", disconnected)
		action = gnet.Shutdown
	}
	return
}

func (s *SimpleServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	codec := c.Context().(*protocol.SimpleCodec)
	var packets [][]byte
	for {
		data, err := codec.Decode(c)
		if err == protocol.ErrIncompletePacket {
			break
		}
		if err != nil {
			zlog.Errorf("invalid packet: %v", err)
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
	return
}
