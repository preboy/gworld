package tcp

import (
	"net"
	"sync"

	"gworld/core/log"
)

type TcpServer struct {
	listener *net.TCPListener
	w        *sync.WaitGroup
	f        func(*net.TCPConn)
}

func NewTcpServer() *TcpServer {
	return &TcpServer{
		w: &sync.WaitGroup{},
	}
}

func (server *TcpServer) Start(host string, f func(*net.TCPConn)) bool {
	server.f = f
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return false
	}
	server.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return false
	}
	go server.rt_accept()
	return true
}

func (server *TcpServer) Stop() {
	if server.listener != nil {
		server.listener.Close()
		server.listener = nil
	}
	server.w.Wait()
}

func (server *TcpServer) rt_accept() {
	server.w.Add(1)
	defer func() {
		server.w.Done()
	}()

	for {
		conn, err := server.listener.AcceptTCP()
		if err != nil {
			log.GetLogger().Println("Error AcceptTCP:", err)
			break
		}
		server.f(conn)
	}
}
