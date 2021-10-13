package main

import (
	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
	coord "github.com/abhilashbss/distributed_coordinator/src/coordinator"
)

func main() {

	logger.InitLogger("/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/log/logs_testing.txt")
	coord1 := coord.CoordActor{}
	coord1.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_1.conf"
	coord1.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	coord1.LoadCoordinator()
	coord1.Listen()

	coord2 := coord.CoordActor{}
	coord1.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_2.conf"
	coord1.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	coord2.LoadCoordinator()
	coord2.Listen()
}
