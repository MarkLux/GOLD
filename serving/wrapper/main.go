package main

import (
	"github.com/MarkLux/GOLD/serving/wrapper/gold"
	"log"
	"os"
	"strconv"
)

func main() {
	s := &gold.GoldService{}
	err := s.LoadComponents()
	if err != nil {
		log.Fatal("fail to launch restful, ", err)
	}
	s.LaunchService()
	serverPid := os.Getpid()
	//println(strconv.Itoa(serverPid))
	funcServicePidFile, err := os.OpenFile("/Users/hanxiao/funcServicePid",os.O_RDWR|os.O_CREATE,0766)
	if err != nil {
		log.Fatal("fail to open funcServicePidFile ")
	}
	defer funcServicePidFile.Close()
	_, err = funcServicePidFile.WriteString(strconv.Itoa(serverPid))
	if err != nil {
		log.Fatal("fail to write pid to file")
	}
}
