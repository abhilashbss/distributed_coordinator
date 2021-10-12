package coordinator

import (
	"encoding/json"
	"fmt"

	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
)

type CoordActor struct {
	Node_count             int                             `json:"Node_count"`
	Seed_node              int                             `json:"Seed_node"`
	Node_listeners         []CommonConfig.Node_url_mapping `json:"Node_listeners"`
	Node_number            int                             `json:"Node_number"`
	Service_specific_data  string                          `json:"Service_specific_data"`
	MsgSender              MessageSender
	Cluster_op_msg_handler MessageHandlerList
}

func (c *CoordActor) RequestHandler(m Message) {

}

// Resigter Actions for each service as per MessageHandlerMapping

// Issue in from node, as it need to store IP instead of node_number

// Messages

func (c *CoordActor) LoadConfMessage() Message {
	var message Message
	message.FromNode = -1
	message.ToNode = c.Seed_node
	message.ServiceName = "coordinator"
	message.ContentData.Action = "New_Node"
	message.ContentData.Data = ""
	return message
}

// Actions

func (c *CoordActor) SendConnectingMsgToSeedNode() {
	msg := c.LoadConfMessage()
	c.MsgSender.MessagePacket = msg
	c.MsgSender.SendMessage()
}

func (c *CoordActor) ExecuteActionForMessage(m Message) {
	c.Cluster_op_msg_handler.ExecuteForAction(m.ContentData.Action, m)
}

// Handlers

func (c *CoordActor) AddMessageHandlers() {

	//New Node request Handler
	var msgHandler MessageHandler
	msgHandler.MessagePacket.ContentData.Action = "New_Node"
	msgHandler.ServiceHandler = c.SendNewNodeResponse
	c.Cluster_op_msg_handler.AddMessageHandler(msgHandler)

	//New Node response Handler
	msgHandler.MessagePacket.ContentData.Action = "New_Node_Response"
	msgHandler.ServiceHandler = c.UpdateCoordMeta
}

// Coordinator Handlers
// vv

type NewNodeCommunicator struct {
	Node_count            int                `json:"Node_count"`
	Node_listeners        []Node_url_mapping `json:"Node_listeners"`
	Service_specific_data string             `json:"Service_specific_data"`
}

func (c *CoordActor) SendNewNodeResponse(m Message) {

	var NewMessage Message
	fromNode := m.FromNode
	for _, node := range c.Node_listeners {
		if node.Node_id == fromNode {
			NewMessage.ContentData.Action = "New_Node_Response"
			NewMessage.FromNode = c.Node_number
			NewMessage.ToNode = fromNode
			NewMessage.ServiceName = "coordinator"
			var newNodeMsg NewNodeCommunicator
			newNodeMsg.Node_count = c.Node_count
			newNodeMsg.Node_listeners = c.Node_listeners
			newNodeMsg.Service_specific_data = c.Service_specific_data

			messageJson, err := json.Marshal(newNodeMsg)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}
			NewMessage.ContentData.Data = string(messageJson)

			c.MsgSender.MessagePacket = NewMessage
			c.MsgSender.SendMessage()
		}
	}
	// for complete new node write the logic here

}

func (c *CoordActor) UpdateCoordMeta(m Message) {
	var msgCommunicator NewNodeCommunicator
	json.Unmarshal([]byte(m.ContentData.Data), &msgCommunicator)
	c.Node_count = msgCommunicator.Node_count
	c.Node_listeners = msgCommunicator.Node_listeners
	c.Service_specific_data = msgCommunicator.Service_specific_data
}

// ^^
//Coordinator Handlers
