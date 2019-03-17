package main

import (
	"github.com/MarkLux/GOLD/serving/wrapper/gold"
	"log"
)

func main() {
	s := &gold.GoldService{}
	s.LoadComponents()
	err := s.LaunchService()
	if err != nil {
		log.Fatal(err)
	}
}
