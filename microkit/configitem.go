package microkit

type ConfigItem struct{
	valueType string
	value string
}

func InitConfigItem (value string, valueType string) ConfigItem {
	return ConfigItem{value: value, valueType: valueType}
}

func (c ConfigItem) GetValue () string {
	return c.value
}

