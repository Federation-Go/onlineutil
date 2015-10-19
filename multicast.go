package multicast

import (
	"errors"
	"net"
)

type GetMessageFunc func() string

type Heartbeater struct {
	msg        string
	group      string
	port       string
	multicast  *net.UDPConn
	getMessage GetMessageFunc
}

func (h *Heartbeater) connect() error {
	addr_string := net.JoinHostPort(h.group, h.port)
	addr, err := net.ResolveUDPAddr("udp", addr_string)
	if err != nil {
		return err
	}
	h.multicast, err = net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	return nil

}
func (h *Heartbeater) beat() error {
	msg := h.msg
	if h.getMessage != nil {
		msg = h.getMessage()
	}
	_, err := h.multicast.Write([]byte(msg))
	if err != nil {
	}
	return err
}
func NewHeartbeater(msg, group, port string, getMessage GetMessageFunc) (*Heartbeater, error) {
	h := new(Heartbeater)
	h.msg = msg
	h.group = group
	h.port = port
	h.getMessage = getMessage
	if err := h.connect(); err != nil {
		return nil, err
	}
	return h, nil
}

type HeartbeatClient struct {
	*Heartbeater
}
