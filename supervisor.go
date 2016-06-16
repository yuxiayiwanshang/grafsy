package main

import (
	"net"
	"os"
	"fmt"
)

type Supervisor struct {
	name string
}

// I want to make this function universal in case we support other daemonization in the future
func (s Supervisor) notify() {
	switch s.name  {
	case "systemd":
		fmt.Println("systemd")
		state := "READY=1"
		socketAddr := &net.UnixAddr{
			Name: os.Getenv("NOTIFY_SOCKET"),
			Net:  "unixgram",
		}

		if socketAddr.Name == "" {
			return
		}

		conn, err := net.DialUnix(socketAddr.Net, nil, socketAddr)
		if err != nil {
			return
		}

		_, err = conn.Write([]byte(state))
		if err != nil {
			return
		}
		fmt.Println("Wrote")
	}

}