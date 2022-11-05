package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"siem-data-producer/producectl/attach_simulator"
	"siem-data-producer/producectl/constants"
	continuesPublisher "siem-data-producer/producectl/continues_producer"
	"siem-data-producer/producectl/log_utils"
	"strings"
)

func init() {
	go constants.UpdateOverrides()
}

func main() {
	serverIP := kingpin.Flag("server", "server address").Required().String()
	protocol := kingpin.Flag("protocol", "server address").Required().Enum("udp", "tcp")

	continues := kingpin.Command("continues", "Produce messages continues_producer")
	attack := kingpin.Command("attack", "Produce messages continues_producer")
	once := kingpin.Command("once", "Produce messages continues_producer")

	// params for continues_producer producer
	filePath := continues.Flag("file_path", "path to read records from").Required().String()
	eps := continues.Flag("eps", "path to read records from").Required().Int()

	// params for attack
	attackFrequencyInSec := attack.Flag("frequency", "How frequent data needs to be produced").Required().Int()
	logLinesAttack := attack.Flag("log_lines", "log lines to be processed").String()
	filePathAttack := attack.Flag("file_path", "path to read records from").String()

	// params for once
	logLines := once.Flag("logLines", "log lines to be processed").Strings()
	filePathOnce := once.Flag("file_path", "path to read records from").String()
	epsOnce := once.Flag("eps", "path to read records from").Int()

	switch kingpin.Parse() {
	case "continues":
		log_utils.Log.Println("Continues producer invoked.")
		continuesPublisher.StartContinuesProducer(*serverIP, *protocol, *filePath, *eps, true)
		break
	case "attack":
		if len(*logLinesAttack) == 0 && *filePathAttack == "" {
			log_utils.Log.Infoln("Log Lines or File Path is needed to run attack")
		}
		log_utils.Log.Infoln("Attack simulator invoked.")
		if *filePathAttack != "" {
			attach_simulator.StartAttackFromFile(*serverIP, *protocol, *filePathAttack, 1, *attackFrequencyInSec)
			break
		}
		attach_simulator.StartAttack(*serverIP, *protocol, strings.Split(*logLinesAttack, ","), *eps, *attackFrequencyInSec)
		break
	case "once":
		if len(*logLines) == 0 && *filePathOnce == "" {
			log_utils.Log.Fatalln("--logLines or --file_path is mandatory. Please use --help for more details")
		}
		log_utils.Log.Infoln("once producer", *serverIP, *protocol, logLines)
		if *epsOnce == 0 {
			*epsOnce = 1
		}
		continuesPublisher.StartContinuesProducer(*serverIP, *protocol, *filePath, *epsOnce, true)
		break
	}

}
