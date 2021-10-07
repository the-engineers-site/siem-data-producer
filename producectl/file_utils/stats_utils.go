package file_utils

import (
	"os"
	"siem-data-producer/producectl/log_utils"
)

func DisplayStats(filePath string) {
	file, err := os.Stat(filePath)
	if err != nil {
		log_utils.Log.Fatalln(err)
	}
	if file.IsDir() {
		log_utils.Log.Fatalln("Specified path is directory. Please specify path till file")
	}
	log_utils.Log.Infoln("File Info")

	fileSizeInMb := (float64)(file.Size() / 1024)
	log_utils.Log.Infoln("File Size", fileSizeInMb)
	if fileSizeInMb < 0.1 {
		log_utils.Log.Infoln(file.Size(), " Bytes")
	} else {
		log_utils.Log.Infoln(fileSizeInMb, " Kb")
	}
}
