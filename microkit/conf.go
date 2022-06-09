package microkit

type Conf struct{
	baseUrl string
}

var instance *Conf = nil

func NewConf(baseUrl string) *Conf {
	if instance == nil {
		instance = new(Conf)
		instance.baseUrl = baseUrl
	}
	return instance
}


