package produce

import (
	"bufio"
	"net"
	"os"
	"siem-data-producer/producectl/log_utils"
	"siem-data-producer/producectl/tcp_utils"
	"time"
)

func PushLogs(connection *net.Conn, filePath string, eps int) {

	file, err := os.Open(filePath)

	if err != nil {
		log_utils.Log.Fatalln(err)
	}

	defer file.Close()

	guard := make(chan struct{}, eps)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		guard <- struct{}{}
		go func(conn *net.Conn, data string) {
			tcp_utils.Publish(conn, data)
			time.Sleep(1 * time.Second)
			<-guard
		}(connection, scanner.Text())
	}

}
