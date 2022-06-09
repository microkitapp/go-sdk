package microkit

import (
	"testing"
	"reflect"
	"fmt"
)

func TestMicroKit(t *testing.T)  {
	jsonv := `{"config":{"db":{"username":{"type": "string", "value": "test"}}}}`
	uri := "3"
	port := 8083
	MockServer(jsonv, port, uri)
	microKit := InitKit("test_key", map[string]string{"url": fmt.Sprintf("http://localhost:%d/%s", port, uri)})
	if reflect.TypeOf(*microKit).String() != "microkit.MicroKit" {
		t.Errorf("got %s want %s", reflect.TypeOf(*microKit), "microkit.MicroKit")
	}
}

func TestMicroKitReady(t *testing.T)  {
	jsonv := `{"config":{"db":{"username":{"type": "string", "value": "test"}}}}`
	uri := "4"
	port := 8084
	MockServer(jsonv, port, uri)
	microKit := InitKit("test_key", map[string]string{"url": fmt.Sprintf("http://localhost:%d/%s", port, uri)})
	
	if !microKit.KitReady  {
		t.Errorf("got %t want %t", microKit.KitReady, true)
	}
}

func TestMicroKitConfigKitType(t *testing.T)  {
	jsonv := `{"config":{"db":{"username":{"type": "string", "value": "test"}}}}`
	uri := "5"
	port := 8085
	MockServer(jsonv, port, uri)
	microKit := InitKit("test_key", map[string]string{"url": fmt.Sprintf("http://localhost:%d/%s", port, uri)})
	
	if reflect.TypeOf(microKit.ConfigKit).String() != "map[string]microkit.ConfigGroup" {
		t.Errorf("got %s want %s", reflect.TypeOf(microKit.ConfigKit), "map[string]microkit.ConfigGroup")
	}
}

func TestMicroKitConfigKitValue(t *testing.T)  {
	jsonv := `{"config":{"db":{"username":{"type": "string", "value": "test"}}}}`
	uri := "6"
	port := 8086
	MockServer(jsonv, port, uri)
	microKit := InitKit("test_key", map[string]string{"url": fmt.Sprintf("http://localhost:%d/%s", port, uri)})
	
	if reflect.TypeOf(microKit.ConfigKit["db"]).String() != "microkit.ConfigGroup" {
		t.Errorf("got %s want %s", reflect.TypeOf(microKit.ConfigKit["db"]), "microkit.ConfigGroup")
	}
}