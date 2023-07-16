package system

import "net"

func Listen(network, address string) (net.Listener, error) {
	if l, e := net.Listen(network, address); e != nil {
		return nil, e
	} else {
		return l, nil
	}
}
