package service

import (
	"fmt"

	messaging "github.com/abhilashbss/distributed_coordinator/src/messaging"
)

type ExampleService struct {
	ServiceObj Service
}

func (e *ExampleService) Handlerping(m messaging.Message) {
	fmt.Println("Ping recieved at : " + e.ServiceObj.Node_addr + " from : " + m.FromNode)
	e.ActionSendPong(m)
}

func (e *ExampleService) Handlerpong(m messaging.Message) {
	fmt.Println("Pong recieved at : " + e.ServiceObj.Node_addr + " from : " + m.FromNode)
}

func (e *ExampleService) ActionSendPong(m messaging.Message) {
	fmt.Println("Pong sending to: " + e.ServiceObj.Node_addr + " from : " + e.ServiceObj.Node_addr)
	m.ToNode = m.FromNode
	m.FromNode = e.ServiceObj.Node_addr
	m.ContentData.Action = "pong"
	e.ServiceObj.MsgSender.MessagePacket = m
	e.ServiceObj.MsgSender.SendMessage()
}

func (e *ExampleService) ActionSendPing(ToNode string) {
	fmt.Println("Ping sending to: " + ToNode + " from : " + e.ServiceObj.Node_addr)
	m := e.ServiceObj.GetMessageObject(ToNode, "ping", "")
	e.ServiceObj.MsgSender.MessagePacket = m
	e.ServiceObj.MsgSender.SendMessage()
}

func (e *ExampleService) AddHandlers() {
	e.ServiceObj.RegisterHandler("ping", e.Handlerping)
	e.ServiceObj.RegisterHandler("pong", e.Handlerpong)
}
