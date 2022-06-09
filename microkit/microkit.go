package microkit

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

type MicroKit struct{
	KitReady bool
	ConfigKit map[string]ConfigGroup
}


var kitInstance *MicroKit = nil

func  InitKit(sdk_key string, options map[string]string) (*MicroKit) {
	if kitInstance == nil {
		kitInstance = new(MicroKit)
		client := new(HttpClient)
		resp := client.Get(options["url"], sdk_key)
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var rawStructJson map[string]map[string]map[string]map[string]string
		json.Unmarshal(bodyBytes, &rawStructJson)
		// log.Fatalf("%v", rawStructJson)
		kitInstance.ConfigKit = InitConfigKit(rawStructJson["config"])
		kitInstance.KitReady = true
	}
	return kitInstance
}

