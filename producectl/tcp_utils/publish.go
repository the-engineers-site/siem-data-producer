package tcp_utils

import (
	"fmt"
	"net"
	"siem-data-producer/producectl/formatter"
	"siem-data-producer/producectl/log_utils"
)

func Publish(connection *net.Conn, logLine string) {
	if connection == nil {
		log_utils.Log.Errorln("Empty connection")
		return
	}
	formattedList := formatter.FormatLog(logLine)
	_, err := fmt.Fprintln(*connection, formattedList)
	if err != nil {
		log_utils.Log.Debugln("Error while publishing ", err)
	}
}

func GetConnection(ip string, protocol string) (net.Conn, error) {
	return net.Dial(protocol, ip)
}
