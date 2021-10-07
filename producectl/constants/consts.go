package constants

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"siem-data-producer/models/configuration"
	"siem-data-producer/producectl/log_utils"
	"time"
)

const (
	OVERRIDE = "/v1/configuration"
)

var Overrides map[string][]string

func UpdateOverrides() {
	if Overrides == nil {
		Overrides = make(map[string][]string)
	}
	apiServer := os.Getenv("API_SERVER")
	if apiServer == "" {
		log_utils.Log.Warnln("Selecting default api server")
		apiServer = "http://localhost:8082"
	}
	fetchDataFromAPI(apiServer)
	for range time.Tick(time.Duration(5) * time.Minute) {
		fetchDataFromAPI(apiServer)
	}

}

func fetchDataFromAPI(apiServer string) {
	resp, err := http.Get(apiServer + OVERRIDE)
	if err != nil {
		log_utils.Log.Errorln("Unable to fetch overrides. Please check connection. ", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log_utils.Log.Errorln(err)
		}
	}(resp.Body)

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log_utils.Log.Errorln("Unable to extract response. ", err)
		return
	}
	var configurations []configuration.Configuration
	err = json.Unmarshal(responseBytes, &configurations)
	if err != nil {
		log_utils.Log.Errorln("Unable to unmarshal log", err)
		return
	}
	log_utils.Log.Infoln("fetched ", len(configurations))

	for _, item := range configurations {
		Overrides[item.OverrideKey] = item.OverrideValues
	}
	log_utils.Log.Infoln("Loading completed")
}
