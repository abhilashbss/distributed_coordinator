package service

import (
	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
	messaging "github.com/abhilashbss/distributed_coordinator/src/messaging"
)

type ServiceGroup struct {
	Node_addr      string                          `json:"Node_addr"`
	Node_listeners []CommonConfig.Node_url_mapping `json:"Node_listeners"`
	ServiceList    []Service                       `json:"ServiceList"`
	MsgSender      messaging.MessageSender
}

func (s *ServiceGroup) InitService(Node_addr string, Node_listeners []CommonConfig.Node_url_mapping, MsgSender messaging.MessageSender) {
	s.Node_addr = Node_addr
	s.Node_listeners = Node_listeners
	s.MsgSender = MsgSender
}

func (s *ServiceGroup) AddService(service Service) {
	service.AddNodeOperations(s.Node_addr, s.Node_listeners, s.MsgSender)
	s.ServiceList = append(s.ServiceList, service)
}

func (s *ServiceGroup) ExecuteMessageAction(msg messaging.Message) {
	for _, service := range s.ServiceList {
		if service.GetServiceName() == msg.ServiceName {
			handler := service.GetServiceHandler()
			handler.ExecuteForAction(msg.ContentData.Action, msg)
		}
	}
}

// create ServiceGroup object in main
// then create a new service as instance of Service and register handlers
// register to ServiceGroup
// Pass the same ServiceGroup object to coord
