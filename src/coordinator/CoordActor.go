package coordinator

import (
	"encoding/json"
	"fmt"

	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
	Util "github.com/abhilashbss/distributed_coordinator/src/util"
)

type CoordActor struct {
	Node_count             int                             `json:"Node_count"`
	Seed_node              int                             `json:"Seed_node"`
	Node_listeners         []CommonConfig.Node_url_mapping `json:"Node_listeners"`
	Node_number            int                             `json:"Node_number"`
	Service_specific_data  string                          `json:"Service_specific_data"`
	Seed_addr              string                          `json:"Seed_address"`
	Node_addr              string                          `json:"Node_address"`
	MsgSender              MessageSender
	Cluster_op_msg_handler MessageHandlerList
}

func (c *CoordActor) RequestHandler(m Message) {

}

func (c *CoordActor) LoadCoordinator() {
	filepath := "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init.conf"
	nodeConf, _ := Util.LoadNodeConf(filepath)

	if nodeConf.Current_node == nodeConf.Seed_node {
		c.LoadSeedCoordinator()
	} else {
		c.LoadNodeCoordinator()
	}
}

// Resigter Actions for each service as per MessageHandlerMapping

// Issue in from node, as it need to store IP instead of node_number

// Messages

func (c *CoordActor) ExecuteActionForMessage(m Message) {
	c.Cluster_op_msg_handler.ExecuteForAction(m.ContentData.Action, m)
}

// Handlers

// Coordinator Handlers
// vv

type NewNodeCommunicator struct {
	Node_count            int                             `json:"Node_count"`
	Node_listeners        []CommonConfig.Node_url_mapping `json:"Node_listeners"`
	Service_specific_data string                          `json:"Service_specific_data"`
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

func (c *CoordActor) AddNewNodeMessageHandlers() {
	var msgHandler MessageHandler
	msgHandler.MessagePacket.ContentData.Action = "New_Node"
	msgHandler.ServiceHandler = c.SendNewNodeResponse
	c.Cluster_op_msg_handler.AddMessageHandler(msgHandler)
}

// ^^
//Coordinator Handlers
