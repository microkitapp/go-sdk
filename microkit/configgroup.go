package microkit

type ConfigGroup struct{
	valueType string
	value map[string]ConfigItem
}

func InitConfigGroup (value map[string]map[string]string) ConfigGroup {
	group := make(map[string]ConfigItem, len(value))
	for k := range value {
		group[k] = InitConfigItem(value[k]["value"], value[k]["type"])
	}
	return ConfigGroup{value: group, valueType: "group"}
}

func (c ConfigGroup) GetValue () (map[string]ConfigItem) {
	return c.value
}

