package single

import "sync"

type singleton struct {
}

type singletonOnce struct {
	sync.Once
	instance *singleton
}

var singletonInstance = &singletonOnce{}

func GetInstance() *singleton {

	singletonInstance.Do(func() {
		singletonInstance.instance = &singleton{}
	})

	return singletonInstance.instance
}
