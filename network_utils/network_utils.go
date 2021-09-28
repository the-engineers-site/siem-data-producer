package network_utils

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/Formatter"
	"gitlab.com/yjagdale/siem-data-producer/models/producer"
	"gitlab.com/yjagdale/siem-data-producer/models/profile"
)

func StartProducer(p *producer.Producer) producer.Response {
	log.Infoln("Starting producer")
	var response producer.Response
	stats, err := os.Stat(p.Profile.FilePath)

	if err != nil {
		response.SetMessage(http.StatusNotFound, nil, err)
		return response
	}

	if stats.Size() > 1594682 {
		log.Infoln("File is greater than 1594682")
		log.Infoln("Producing aync")
		p.Continues = true
	}

	if p.Continues {
		go func() {
			err := readAndPushLogsAsync(p.Profile, p.Eps)
			if err != nil {
				log.Errorln("Error while producing logs.", err)
			}
		}()
	} else {
		err := readAndPushLogsAsync(p.Profile, p.Eps)
		if err != nil {
			response.SetMessage(http.StatusInternalServerError, nil, err)
			return response
		}
	}
	response.SetMessage(http.StatusAccepted, gin.H{"execution_id": p.ExecutionId}, nil)
	return response
}

func readAndPushLogsAsync(profile *profile.Profile, eps int) error {
	log.Infoln("Async producer started")
	file := readFile(profile.FilePath)
	log.Infoln("File read completed")
	connection, err := getConnection(profile.Destination, profile.Protocol)
	log.Infoln("Connection eastablished")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error("Error while closing file", err.Error())
		}
	}(file)

	scanner := bufio.NewScanner(file)

	guard := make(chan struct{}, eps)
	log.Infoln("Posting logs to destination")
	for scanner.Scan() {
		guard <- struct{}{}
		go func(conn net.Conn, data string) {
			produceLog(conn, data)
			time.Sleep(1 * time.Second)
			<-guard
		}(connection, scanner.Text())
	}

	return nil

}

func produceLog(connection net.Conn, log string) bool {
	logLine := Formatter.FormatLog(log)
	err := pushLog(connection, logLine)
	if err != nil {
		return false
	} else {
		return true
	}
}

func pushLog(connection net.Conn, logLine string) error {
	log.Debugln(logLine)
	noOfBytes, err := fmt.Fprintln(connection, logLine)
	if err != nil {
		return err
	}

	log.Debugln("Published ", noOfBytes)
	return nil
}

func getConnection(destinationServer string, protocol string) (net.Conn, error) {
	var conn net.Conn
	var err error
	switch strings.ToUpper(protocol) {
	case "TCP":
		log.Infoln("Building tcp connection")
		conn, err = net.DialTimeout("tcp", destinationServer, 40*time.Second)
		break
	case "UDP":
		log.Infoln("Building UDP connection")
		conn, err = net.Dial("udp", destinationServer)
		break
	}
	if err != nil {
		log.Errorln("could not connect to server: ", err)
		return nil, err
	}
	return conn, nil
}

func readFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Error(err)
	}
	return file
}
