package main

import (
	"github.com/MarkLux/GOLD/serving/wrapper/gold"
)

func main() {
	s := &gold.GoldService{}
	s.LoadComponents()
	s.LaunchService()
}
