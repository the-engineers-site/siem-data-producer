package attach_simulator

import (
	"net"
	"siem-data-producer/producectl/file_utils"
	"siem-data-producer/producectl/log_utils"
	"siem-data-producer/producectl/produce"
	"siem-data-producer/producectl/tcp_utils"
	"time"
)

func StartAttack(serverIp string, protocol string, logLines []string, eps int, duration int) {
	log_utils.Log.Debugln("EPS", eps)
	connection, err := getConnection(serverIp, protocol)
	if err != nil {
		log_utils.Log.Fatalln(err)
	}

	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			log_utils.Log.Errorln(err)
		}
	}(connection)
	for range time.Tick(time.Duration(duration) * time.Second) {
		for _, line := range logLines {
			tcp_utils.Publish(&connection, line)
			time.Sleep(1 * time.Second)
		}
	}
}

func StartAttackFromFile(serverIp string, protocol string, filePath string, eps int, duration int) {
	file_utils.DisplayStats(filePath)
	log_utils.Log.Infof("Destination: %s", serverIp)
	log_utils.Log.Infof("Protocol: %s", protocol)
	log_utils.Log.Infof("eps: %v", eps)
	log_utils.Log.Infoln("Connection opened")
	connection, err := getConnection(serverIp, protocol)
	if err != nil {
		log_utils.Log.Fatalln(err)
	}

	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			log_utils.Log.Errorln(err)
		}
	}(connection)

	for range time.Tick(time.Duration(duration) * time.Second) {
		produce.PushLogs(&connection, filePath, eps)
		log_utils.Log.Infoln("Attack published for ", filePath)
	}
}

func getConnection(serverIp string, protocol string) (net.Conn, error) {
	return tcp_utils.GetConnection(serverIp, protocol)
}
