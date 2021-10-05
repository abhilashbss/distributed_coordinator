package coordinator

type Message struct {
	ServiceName string `json: "ServiceName"`
	FromNode    int    `json: "FromNode"`
	ToNode      int    `json: "ToNode"`
	Content     string `json: "Content"`
}

type MessageHandler struct {
	MessagePacket  Message       `json: "message"`
	ServiceHandler func(Message) `json: "ServiceHandler"`
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
