package common

import (
	"os"
	"sync"
)

type GoldEnv struct {
	PodName string
}

var gENV *GoldEnv = nil
var envOnce sync.Once

func InitGoldEnv() error {
	hostName, err := os.Hostname()
	if err != nil {
		return err
	}
	gENV = &GoldEnv{PodName: hostName}
	return nil
}

func GetGoldEnv() (env *GoldEnv, err error) {
	if gENV != nil {
		env = gENV
	} else {
		envOnce.Do(func() {
			err = InitGoldEnv()
			env = gENV
		})
	}
	return
}
