package coordinator

type Content struct {
	Action string `json: "Action"`
	Data   string `json:Data`
}

type Message struct {
	ServiceName string  `json: "ServiceName"`
	FromNode    int     `json: "FromNode"`
	ToNode      int     `json: "ToNode"`
	ContentData Content `json: "Content"`
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

type MessageHandlerList struct {
	MessageHandlerList []MessageHandler
}

func (m *MessageHandlerList) FindHandlerForAction(Action string) MessageHandler {
	for _, mh := range m.MessageHandlerList {
		if mh.MessagePacket.ContentData.Data == Action {
			return mh
		}
	}
	return MessageHandler{}
}

//Initally register Message Handler with dummy message for Action mapping
func (m *MessageHandlerList) ExecuteForAction(Action string) {
	mh := m.FindHandlerForAction(Action)
	mh.ExecuteMessageHandler()
}

func (m *MessageHandlerList) AddMessageHandler(mh MessageHandler) {
	m.MessageHandlerList = append(m.MessageHandlerList, mh)
}

func (m *MessageHandlerList) ExecuteMessageHandlers() {
	for _, mh := range m.MessageHandlerList {
		mh.ExecuteMessageHandler()
	}
}
