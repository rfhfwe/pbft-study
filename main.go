package main

import (
	"os"

	"github.com/bigpicturelabsinc/consensusPBFT/pbft/network"
)

func main() {
	nodeID := os.Args[1] // 接收一个来自命令行的参数作为节点 Id
	server := network.NewServer(nodeID)

	server.Start()
}
