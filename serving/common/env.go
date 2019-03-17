package common

import (
	"github.com/MarkLux/GOLD/serving/wrapper/constant"
	"os"
	"sync"
)

type GoldEnv struct {
	PodName string
	ServiceName string
}

var gENV *GoldEnv = nil
var envOnce sync.Once

func InitGoldEnv() error {
	hostName, err := os.Hostname()
	if err != nil {
		return err
	}
	serviceName := os.Getenv(constant.GoldServiceNameEnvKey)
	gENV = &GoldEnv{PodName: hostName, ServiceName: serviceName}
	return nil
}

func GetGoldEnv() (env *GoldEnv) {
	if gENV != nil {
		env = gENV
	} else {
		envOnce.Do(func() {
			err := InitGoldEnv()
			if err != nil {
				env = &GoldEnv{ServiceName: "Unknown", PodName: "Unknown"}
			} else {
				env = gENV
			}
		})
	}
	return
}
