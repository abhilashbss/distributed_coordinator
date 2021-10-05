package coordinator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MessageType int

const (
	Coordinator MessageType = iota
	Service
)

type Node_url_mapping struct {
	Node_id          int
	Node_listner_url string
}

// Usability : set message, messageType and init Node Listeners and then send message
type MessageSender struct {
	MessagePacket   Message
	Type_of_message MessageType
	Node_listeners  []Node_url_mapping
}

func (m *MessageSender) FindListener(Node_id int) string {
	for _, node_url := range m.Node_listeners {
		if node_url.Node_id == Node_id {
			return node_url.Node_listner_url
		}
	}
	return ""
}

// setters and getters
// update the listener address of other nodes

func (m *MessageSender) SendMessage() {
	ToNode := m.MessagePacket.ToNode
	listener := m.FindListener(ToNode)

	messageJson, err := json.Marshal(m.MessagePacket)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	var jsonStr = []byte(messageJson)
	var msgType string
	if m.Type_of_message == Coordinator {
		msgType = "/coordinator_request"
	} else {
		msgType = "/service_request"
	}
	req, err := http.NewRequest("POST", listener+msgType, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Non-OK HTTP status:", resp.StatusCode)
		return
	}
}

func (m *MessageSender) SetMessagePacket(msg Message) {
	m.MessagePacket = msg
}

func (m *MessageSender) SetMessageType(msgType MessageType) {
	m.Type_of_message = msgType
}

func (m *MessageSender) SetNodeListeners(url_mapping []Node_url_mapping) {
	m.Node_listeners = url_mapping
}
