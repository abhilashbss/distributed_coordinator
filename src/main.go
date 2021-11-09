package main

import (
	"sync"
	"time"

	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
	coord "github.com/abhilashbss/distributed_coordinator/src/coordinator"
)

func main() {
	logger.InitLogger("/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/log/logs.txt")
	coord1 := coord.CoordActor{}
	coord1.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init.conf"
	coord1.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	coord1.LoadCoordinator()
	coord1.Listen()
}

func FourNodeSeedNodeJoinTest() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(4)
	// Seed node is run first
	logger.InitLogger("/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/log/logs.txt")
	coord1 := coord.CoordActor{}
	coord1.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_1.conf"
	coord1.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	coord1.LoadCoordinator()
	go func() {
		coord1.Listen()
	}()
	time.Sleep(1 * time.Second)

	// All other nodes
	coord2 := coord.CoordActor{}
	coord2.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_2.conf"
	coord2.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	go func() {
		coord2.Listen()
	}()
	time.Sleep(1 * time.Second)
	coord2.LoadCoordinator()

	coord3 := coord.CoordActor{}
	coord3.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_3.conf"
	coord3.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	go func() {
		coord3.Listen()
	}()
	time.Sleep(1 * time.Second)
	coord3.LoadCoordinator()

	coord4 := coord.CoordActor{}
	coord4.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_4.conf"
	coord4.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	go func() {
		coord4.Listen()
	}()
	time.Sleep(1 * time.Second)
	coord4.LoadCoordinator()

	waitGroup.Wait()
}
