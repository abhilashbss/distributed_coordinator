package Testing

import (
	"sync"
	"testing"

	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
	coord "github.com/abhilashbss/distributed_coordinator/src/coordinator"
	service "github.com/abhilashbss/distributed_coordinator/src/service"
)

func ServiceCommunication(t *testing.T) {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	var serviceGroup1 service.ServiceGroup
	logger.InitLogger("/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/log/logs.txt")
	coord1 := coord.CoordActor{}
	coord1.Service_message_processor = serviceGroup1
	coord1.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_1.conf"
	coord1.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	go func() {
		coord1.Listen()
	}()
	coord1.LoadCoordinator()

	var serviceGroup2 service.ServiceGroup
	coord2 := coord.CoordActor{}
	coord2.Service_message_processor = serviceGroup2
	coord2.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_2.conf"
	coord2.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	go func() {
		coord2.Listen()
	}()
	coord2.LoadCoordinator()

	var exService1 service.ExampleService
	exService1.AddHandlers()
	exService1.RegisterToServiceGroup(coord1.Service_message_processor)
	exService1.ActionSendPing("localhost:8082")

	var exService2 service.ExampleService
	exService2.AddHandlers()
	exService2.RegisterToServiceGroup(coord2.Service_message_processor)

	waitGroup.Wait()
}
