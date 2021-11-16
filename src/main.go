package main

import (
	"flag"
	"fmt"
	"sync"

	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
	coord "github.com/abhilashbss/distributed_coordinator/src/coordinator"
	service "github.com/abhilashbss/distributed_coordinator/src/service"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	arg_map := parseArgs()
	serviceGroup := InitServices()
	fmt.Println(arg_map)
	logger.InitLogger(arg_map["log_path"])
	coord1 := coord.CoordActor{}
	coord1.Service_message_processor = serviceGroup
	coord1.Node_conf_path = arg_map["node_conf_path"]
	coord1.Cluster_conf_path = arg_map["cluster_meta_path"]
	go func() {
		coord1.Listen()
	}()
	coord1.LoadCoordinator()
	waitGroup.Wait()
}

func parseArgs() map[string]string {
	arg_map := make(map[string]string)

	var log_path, node_init_path, cluster_meta_path string
	flag.StringVar(&log_path, "log_path", "/log", "log file path")
	flag.StringVar(&node_init_path, "node_conf_path", "/node_init.conf", "init conf file path")
	flag.StringVar(&cluster_meta_path, "cluster_conf_path", "/cluster_meta.conf", "cluster meta conf file path")
	flag.Parse() 
	arg_map["log_path"] = log_path
	arg_map["node_conf_path"] = node_init_path
	arg_map["cluster_meta_path"] = cluster_meta_path
	return arg_map
}

func InitServices() service.ServiceGroup {
	var serviceGroup service.ServiceGroup
	var PingPongService service.ExampleService
	PingPongService.Init()
	PingPongService.RegisterToServiceGroup(serviceGroup)
	return serviceGroup
}
