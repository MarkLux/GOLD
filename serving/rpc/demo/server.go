package main

import (
	goldrpc "goldrpc"
)

func main() {
	s := goldrpc.Server()
	s.Serve(":8099")
}