package coordinator

import (
	Util "github.com/abhilashbss/distributed_coordinator/src/util"
)

type CoordActor struct {
	Node_count             int                `json:"Node_count"`
	Seed_node              int                `json:"Seed_node"`
	Node_listeners         []Node_url_mapping `json:"Node_listeners"`
	Node_number            int                `json:"Node_number"`
	Service_specific_data  string             `json:"Service_specific_data"`
	MsgSender              MessageSender
	Cluster_op_msg_handler MessageHandlerList
}

func (c *CoordActor) RequestHandler(m Message) {

}

func (c *CoordActor) LoadCoordinator() {
	filepath := "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	configuration := Util.LoadConf(filepath)

	if configuration.Node_number == configuration.Seed_node {
		c.Node_count = configuration.Node_count
		c.Node_listeners = configuration.Node_listeners
		c.Node_number = configuration.Node_number
		c.Seed_node = configuration.Seed_node
		c.Service_specific_data = configuration.Service_specific_data
	} else {
		c.Seed_node = configuration.Seed_node
		c.Node_listeners = configuration.Node_listeners
		c.MsgSender.SetNodeListeners(configuration.Node_listeners)
		c.SendConnectingMsgToSeedNode()
	}
}

//Resigter Actions for each service as per MessageHandlerMapping

func (c *CoordActor) LoadConfMessage() Message {
	var message Message
	message.FromNode = -1
	message.ToNode = c.Seed_node
	message.ServiceName = "coordinator"
	message.ContentData.Action = "New_Node"
	message.ContentData.Data = ""
	return message
}

func (c *CoordActor) SendConnectingMsgToSeedNode() {
	msg := c.LoadConfMessage()
	c.MsgSender.MessagePacket = msg
	c.MsgSender.SendMessage()
}

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
	Node_count             int                `json:"Node_count"`
	Node_listeners         []Node_url_mapping `json:"Node_listeners"`
	Service_specific_data  string             `json:"Service_specific_data"`
}

func (c *CoordActor) SendNewNodeResponse(m Message) {
	
	var NewMessage Message
	fromNode := m.FromNode
	for _, node := range c.Node_listeners {
		if(node.Node_id == fromNode)
		{
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
