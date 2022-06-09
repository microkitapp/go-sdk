package microkit

import (
	"testing"
)


func TestConfHaveValueProp(t *testing.T)  {
	item := Conf{baseUrl: "test"}
	if item.baseUrl != "test" {
		t.Errorf("got %s want %s", item.baseUrl, "test")
	}
}


func TestNewConf(t *testing.T)  {
	item := NewConf("demo_url")
	if item.baseUrl != "demo_url" {
		t.Errorf("got %s want %s", item.baseUrl, "demo_url")
	}
}

// func TestConfGetValue(t *testing.T)  {
// 	item := Conf{value: "test"}
// 	if item.GetValue() != "test" {
// 		t.Errorf("got %s want %s", item.GetValue() , "test")
// 	}
// }

// func TestConfInit(t *testing.T)  {
// 	item := InitConf("test", "string")
// 	if item.GetValue() != "test" {
// 		t.Errorf("got %s want %s", item.GetValue() , "test")
// 	}
// }