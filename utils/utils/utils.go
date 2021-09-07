package utils

import (
	"fmt"
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

func GetOutboundIP() string {

	if os.Getenv("HOST_IP") != "" {
		return os.Getenv("HOST_IP")
	}

	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return fmt.Sprintf("%s", localAddr.IP)
}
