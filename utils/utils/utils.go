package utils

import (
	log "github.com/sirupsen/logrus"
	"net"
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	return port
}

func GetOutboundIP() net.IP {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
