package microkit

import (
	"testing"
	"reflect"
)


func TestConfigGroupValuePropIsmapString(t *testing.T)  {
	item := ConfigGroup{}
	if reflect.TypeOf(item.value).String() != "map[string]microkit.ConfigItem" {
		t.Errorf("got %s want %s", reflect.TypeOf(item.value), "map[string]microkit.ConfigItem")
	}
}

func TestConfigGroupHaveTypeGroup(t *testing.T)  {
	item := ConfigGroup{valueType: "group"}
	if item.valueType != "group" {
		t.Errorf("got %s want %s", item.valueType, "group")
	}
}

func TestConfigGroupGetValue(t *testing.T)  {
	group := map[string]ConfigItem{ "name": ConfigItem{ }}
	item := ConfigGroup{value: group}
	value := item.GetValue()
	if reflect.TypeOf(value).String() != "map[string]microkit.ConfigItem" {
		t.Errorf("got %v want %v", item.GetValue() , group)
	}
}

func TestConfigGroupInit(t *testing.T)  {
	group := map[string]map[string]string{ "name": {}}
	item := InitConfigGroup(group)
	value := item.GetValue()
	if reflect.TypeOf(value).String() != "map[string]microkit.ConfigItem" {
		t.Errorf("got %v want %v", item.GetValue() , group)
	}
}