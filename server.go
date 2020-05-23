package main

import (
	"fmt"
	"net"
	"runtime"
)

func run(delay int, port int) {
	addr := &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: port,
	}
	c, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer c.Close()
	fmt.Printf("[*] Listening on port %d\n", port)

	errs := make(chan error, 1)
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		go listen(c, errs)
	}

	err = <-errs
	fmt.Print(err)
}

func listen(c *net.UDPConn, errs chan<- error) {
	buf := make([]byte, 1024)
	err := error(nil)

	for {
		_, addr, err := c.ReadFromUDP(buf)
		if err != nil {
			break
		}
		fmt.Printf("[*] Connection from %s\n", addr.String())
		data, err := getResponse(buf)
		if err != nil {
			break
		}
		_, err = c.WriteToUDP(data, addr)
		if err != nil {
			break
		}
	}
	errs <- err
}

func getResponse(packet []byte) ([]byte, error) {
	h, err := parseHeader(packet)
	if err != nil {
		return nil, err
	}

	r := h.getResponse()
	return r.toBinary(), nil
}
