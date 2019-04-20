package main

import (
	"github.com/MarkLux/GOLD/serving/wrapper/gold"
	"log"
)

func main() {
	s := &gold.GoldService{}
	err := s.LoadComponents()
	if err != nil {
		log.Fatal("fail to launch restful, ", err)
	}
	s.LaunchService()
}
