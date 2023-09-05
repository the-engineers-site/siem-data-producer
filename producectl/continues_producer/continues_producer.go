package continues_publisher

import (
	"siem-data-producer/producectl/file_utils"
	"siem-data-producer/producectl/log_utils"
	"siem-data-producer/producectl/produce"
	"siem-data-producer/producectl/tcp_utils"
)

func StartContinuesProducer(serverIp string, protocol string, filePath string, eps int, continues bool) {
	file_utils.DisplayStats(filePath)
	log_utils.Log.Infof("Destination: %s", serverIp)
	log_utils.Log.Infof("Protocol: %s", protocol)
	log_utils.Log.Infof("eps: %v", eps)
	connection, err := tcp_utils.GetConnection(serverIp, protocol)
	if err != nil {
		log_utils.Log.Fatalln(err)
	}
	log_utils.Log.Infoln("Connection opened")
	defer connection.Close()
	if continues {
		produce.PushLogs(&connection, filePath, eps)
	} else {
		produce.PublicLogsOnce(&connection, filePath, eps)
	}
}
