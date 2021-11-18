package service

import (
	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
	messaging "github.com/abhilashbss/distributed_coordinator/src/messaging"
)

type Service interface {
	AddHandlers()
	AddNodeOperations(Node_addr string, Node_listeners []CommonConfig.Node_url_mapping, MsgSender messaging.MessageSender)
	GetServiceName() string
	GetServiceHandler() messaging.MessageHandlerGroup
}

type ServiceOperations struct {
	Service_name    string                          `json:"Service_name"`
	Node_listeners  []CommonConfig.Node_url_mapping `json:"Node_listeners"`
	Node_addr       string                          `json:"Node_addr"`
	Service_handler messaging.MessageHandlerGroup
	MsgSender       messaging.MessageSender
}

type HandlerFunc func(messaging.Message)

func (s *ServiceOperations) RegisterHandler(Action string, Handler HandlerFunc) {
	var msgHandler messaging.MessageHandler
	msgHandler.MessagePacket.ContentData.Action = Action
	msgHandler.ServiceHandler = Handler
	s.Service_handler.AddMessageHandler(msgHandler)
}

func (s *ServiceOperations) GetMessageObject(ToNode string, Action string, ToServiceData string) messaging.Message {
	var message messaging.Message
	message.FromNode = s.Node_addr
	message.ToNode = ToNode
	message.ContentData.Action = Action
	message.ContentData.Data = ToServiceData
	message.ServiceName = s.Service_name
	return message
}

func (s *ServiceOperations) ExecuteAction(Action string, msg messaging.Message) {
	s.Service_handler.ExecuteForAction(Action, msg)
}
