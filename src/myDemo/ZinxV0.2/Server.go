package main

import (
	"Game_Zinx/src/zinx/znet"
)

func main() {
	s := znet.NewServer("[zinx V0.2]")
	s.Serve()
}
