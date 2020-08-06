package client_socks5

import (
	"net"
	"time"

	"github.com/thinkgos/go-socks5/bufferpool"
	"github.com/thinkgos/go-socks5/statute"
)

type underUDPConn struct {
	udpConn       *net.UDPConn
	bufferPool    bufferpool.BufPool
	remoteAddress net.Addr
}

func (sf *underUDPConn) Read(b []byte) (int, error) {
	b1 := sf.bufferPool.Get()
	defer sf.bufferPool.Put(b1)

	n, err := sf.udpConn.Read(b1[:cap(b1)])
	if err != nil {
		return 0, err
	}
	datagram, err := statute.ParseDatagram(b1[:n])
	if err != nil {
		return 0, err
	}
	n = copy(b, datagram.Data)
	return n, nil
}

func (sf *underUDPConn) Write(b []byte) (int, error) {
	datagram, err := statute.NewDatagram(sf.remoteAddress.String(), b)
	if err != nil {
		return 0, err
	}
	return sf.udpConn.Write(datagram.Bytes())
}

func (sf *underUDPConn) Close() error {
	return sf.udpConn.Close()
}

func (sf *underUDPConn) LocalAddr() net.Addr {
	return sf.udpConn.LocalAddr()
}

func (sf *underUDPConn) RemoteAddr() net.Addr {
	return sf.remoteAddress
}

func (sf *underUDPConn) SetDeadline(t time.Time) error {
	return sf.udpConn.SetDeadline(t)
}

func (sf *underUDPConn) SetReadDeadline(t time.Time) error {
	return sf.udpConn.SetReadDeadline(t)
}

func (sf *underUDPConn) SetWriteDeadline(t time.Time) error {
	return sf.udpConn.SetWriteDeadline(t)
}