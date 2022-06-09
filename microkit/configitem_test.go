package microkit

import (
	"testing"
)


func TestConfigItemHaveValueProp(t *testing.T)  {
	item := ConfigItem{value: "test"}
	if item.value != "test" {
		t.Errorf("got %s want %s", item.value, "test")
	}
}

func TestConfigItemHaveTypeProp(t *testing.T)  {
	item := ConfigItem{valueType: "test"}
	if item.valueType != "test" {
		t.Errorf("got %s want %s", item.valueType, "test")
	}
}

func TestConfigItemGetValue(t *testing.T)  {
	item := ConfigItem{value: "test"}
	if item.GetValue() != "test" {
		t.Errorf("got %s want %s", item.GetValue() , "test")
	}
}

func TestConfigItemInit(t *testing.T)  {
	item := InitConfigItem("test", "string")
	if item.GetValue() != "test" {
		t.Errorf("got %s want %s", item.GetValue() , "test")
	}
}