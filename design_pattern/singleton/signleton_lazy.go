package singleton

import "sync"

var (
	lazySingleton *singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &singleton{}
		})
	}
	return lazySingleton
}
