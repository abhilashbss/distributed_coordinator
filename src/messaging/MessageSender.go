package messaging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
)

// Usability : set message, messageType and init Node Listeners and then send message
// after setting Node_listner the first time
// each time just set MessagePacket, Type_of_message and send the message
type MessageSender struct {
	MessagePacket  Message
	Node_listeners []CommonConfig.Node_url_mapping
}

// setters and getters
// update the listener address of other nodes

func (m *MessageSender) SendMessage() {
	listener := m.MessagePacket.ToNode
	messageJson, err := json.Marshal(m.MessagePacket)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	var jsonStr = []byte(messageJson)
	var msgType string
	if m.MessagePacket.ServiceName == "coordinator" {
		msgType = "/coordinator_request"
	} else {
		msgType = "/service_request"
	}
	req, _ := http.NewRequest("POST", "http://"+listener+msgType, bytes.NewBuffer(jsonStr))
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

	logger.InfoLogger.Println("Sent Message : " + string(messageJson))
}

func (m *MessageSender) SetMessagePacket(msg Message) {
	m.MessagePacket = msg
}

func (m *MessageSender) SetNodeListeners(url_mapping []CommonConfig.Node_url_mapping) {
	m.Node_listeners = url_mapping
}
