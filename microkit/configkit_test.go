package microkit

import (
	"testing"
	"encoding/json"
	"reflect"
)


func TestInitConfigKit(t *testing.T)  {
	jsonv := `{"db":{"username":{"type": "string", "value": "test"}}}`
	var rawConfigKit map[string]map[string]map[string]string
	json.Unmarshal([]byte(jsonv), &rawConfigKit)
	
	config := InitConfigKit(rawConfigKit)
	if reflect.TypeOf(config["db"]).String() != "microkit.ConfigGroup"  {
		t.Errorf("got %v want %s", reflect.TypeOf(config["db"]).String(), "microkit.ConfigGroup")
	}
}

func TestConfigKitReady(t *testing.T)  {
	jsonv := `{"db":{"username":{"type": "string", "value": "test"}}}`
	var rawConfigKit map[string]map[string]map[string]string
	json.Unmarshal([]byte(jsonv), &rawConfigKit)
	
	config := InitConfigKit(rawConfigKit)
	if reflect.TypeOf(config["db"]).String() != "microkit.ConfigGroup"  {
		t.Errorf("got %v want %s", reflect.TypeOf(config["db"]).String(), "microkit.ConfigGroup")
	}
}

