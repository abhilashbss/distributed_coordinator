package service

import (
	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
	messaging "github.com/abhilashbss/distributed_coordinator/src/messaging"
)

type ExampleService struct {
	ServiceObj ServiceOperations
}

func (e *ExampleService) Handlerping(m messaging.Message) {
	logger.InfoLogger.Println("Ping recieved at : " + e.ServiceObj.Node_addr + " from : " + m.FromNode)
	e.ActionSendPong(m)
}

func (e *ExampleService) Handlerpong(m messaging.Message) {
	logger.InfoLogger.Println("Pong recieved at : " + e.ServiceObj.Node_addr + " from : " + m.FromNode)
}

func (e *ExampleService) ActionSendPong(m messaging.Message) {
	logger.InfoLogger.Println("Pong sending to: " + e.ServiceObj.Node_addr + " from : " + e.ServiceObj.Node_addr)
	m.ToNode = m.FromNode
	m.FromNode = e.ServiceObj.Node_addr
	m.ContentData.Action = "pong"
	e.ServiceObj.MsgSender.MessagePacket = m
	e.ServiceObj.MsgSender.SendMessage()
}

func (e *ExampleService) ActionSendPing(ToNode string) {
	logger.InfoLogger.Println("Ping sending to: " + ToNode + " from : " + e.ServiceObj.Node_addr)
	m := e.ServiceObj.GetMessageObject(ToNode, "ping", "")
	e.ServiceObj.MsgSender.MessagePacket = m
	e.ServiceObj.MsgSender.SendMessage()
}

func (e *ExampleService) AddHandlers() {
	e.ServiceObj.RegisterHandler("ping", e.Handlerping)
	e.ServiceObj.RegisterHandler("pong", e.Handlerpong)
}

// func (e *ExampleService) RegisterToServiceGroup(serviceGroup ServiceGroup) {
// 	serviceGroup.AddService(e.ServiceObj)
// }

func (e *ExampleService) AddNodeOperations(Node_addr string, Node_listeners []CommonConfig.Node_url_mapping, MsgSender messaging.MessageSender) {
	logger.InfoLogger.Println("Service :" + Node_addr)
	e.ServiceObj.Node_addr = Node_addr
	e.ServiceObj.Node_listeners = Node_listeners
	e.ServiceObj.MsgSender = MsgSender
}

func (e *ExampleService) GetServiceName() string {
	return e.ServiceObj.Service_name
}

func (e *ExampleService) SetServiceName(name string) {
	e.ServiceObj.Service_name = name
}

func (e *ExampleService) GetServiceHandler() messaging.MessageHandlerGroup {
	return e.ServiceObj.Service_handler
}
