package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Subcommands
	continueCommand := flag.NewFlagSet("continues", flag.ExitOnError)
	// 	listCommand := flag.NewFlagSet("once", flag.ExitOnError)

	// Count subcommand flag pointers
	// Adding a new choice for --metric of 'substring' and a new --substring flag
	continuesFilePathPtr := continueCommand.String("file", "", "File to read logs from. (Required)")
	continuesEPSPtr := continueCommand.Int("eps", 1, "EPS through which logs needs to be posted. (Required)")
	continueHostPtr := continueCommand.String("host", "", "Destination Server to produce logs. Required for --metric=substring")

	log.Info(&continuesFilePathPtr, &continuesEPSPtr, continueHostPtr)
}
