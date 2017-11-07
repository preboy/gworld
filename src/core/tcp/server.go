package tcp

import (
	"fmt"
	"net"
	"sync"
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
	}
	server.w.Wait()
}

func (server *TcpServer) rt_accept() {
	defer func() {
		server.w.Done()
	}()
	server.w.Add(1)
	for {
		conn, err := server.listener.AcceptTCP()
		if err != nil {
			fmt.Println("Error AcceptTCP:", err)
			break
		}
		server.f(conn)
	}
}
