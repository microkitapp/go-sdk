package microkit

type Config interface{
	GetValue() interface{}
}

func InitConfigKit(c map[string]map[string]map[string]string) map[string]ConfigGroup {
	groups := make(map[string]ConfigGroup, len(c))
	for k := range c {
		groups[k] = InitConfigGroup(c[k])
	}
	return groups
	 
}


