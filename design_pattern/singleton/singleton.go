package singleton

type singleton struct {
}

var singletonInstance *singleton

func init() {
	singletonInstance = &singleton{}
}

func GetInstance() *singleton {
	return singletonInstance
}
