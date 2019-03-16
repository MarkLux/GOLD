package main

import (
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
)

func main() {
	s := &goldrpc.Server{}
	s.Serve(":8099")
}