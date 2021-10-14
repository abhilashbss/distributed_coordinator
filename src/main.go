package main

import (
	"sync"

	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
	coord "github.com/abhilashbss/distributed_coordinator/src/coordinator"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	logger.InitLogger("/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/log/logs.txt")
	coord1 := coord.CoordActor{}
	coord1.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_1.conf"
	coord1.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	coord1.LoadCoordinator()
	go func() {
		coord1.Listen()
	}()

	coord2 := coord.CoordActor{}
	coord2.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_2.conf"
	coord2.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	go func() {
		coord2.Listen()
	}()
	coord2.LoadCoordinator()

	waitGroup.Wait()
}
