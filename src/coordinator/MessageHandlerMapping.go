package coordinator

import (
	"encoding/json"

	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
)

type Content struct {
	Action string `json:"Action"`
	Data   string `json:"Data"`
}

type Message struct {
	ServiceName string  `json:"ServiceName"`
	FromNode    string  `json:"FromNode"`
	ToNode      string  `json:"ToNode"`
	ContentData Content `json:"Content"`
}

type MessageHandler struct {
	MessagePacket  Message       `json:"message"`
	ServiceHandler func(Message) `json:"ServiceHandler"`
}

func (m *MessageHandler) SetMessagePacket(msg Message) {
	m.MessagePacket = msg
}

func (m *MessageHandler) SetServiceHandler(handler func(Message)) {
	m.ServiceHandler = handler
}

func (m *MessageHandler) ExecuteMessageHandler() {
	m.ServiceHandler(m.MessagePacket)
}

type MessageHandlerList struct {
	MessageHandlerList []MessageHandler
}

func (m *MessageHandlerList) FindHandlerForAction(Action string) MessageHandler {
	for _, mh := range m.MessageHandlerList {
		if mh.MessagePacket.ContentData.Action == Action {
			return mh
		}
	}
	return MessageHandler{}
}

// Initally register Message Handler with dummy message for Action mapping
// For new message, find handler from dummy added msg, then take action with passing the new message
func (m *MessageHandlerList) ExecuteForAction(Action string, msg Message) {
	messageJson, _ := json.Marshal(msg)
	logger.InfoLogger.Println("Executing for msg : " + string(messageJson))
	mh := m.FindHandlerForAction(Action)
	handler := mh.ServiceHandler
	handler(msg)
}

func (m *MessageHandlerList) AddMessageHandler(mh MessageHandler) {
	m.MessageHandlerList = append(m.MessageHandlerList, mh)
}

func (m *MessageHandlerList) ExecuteMessageHandlers() {
	for _, mh := range m.MessageHandlerList {
		mh.ExecuteMessageHandler()
	}
}
