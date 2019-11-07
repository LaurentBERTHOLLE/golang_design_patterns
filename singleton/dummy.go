package singleton

var instance *Dummy

type Dummy struct {
}

func GetInstance() *Dummy {
	if instance == nil {
		instance = &Dummy{}
	}
	return instance
}
