package coordinator

import (
	"encoding/json"
	"fmt"

	Util "github.com/abhilashbss/distributed_coordinator/src/util"
)

func (c *CoordActor) LoadSeedCoordinator() {
	filepath := c.Cluster_conf_path
	configuration, _ := Util.LoadSeedConf(filepath)
	c.Node_count = configuration.Node_count
	c.Node_listeners = configuration.Node_listeners
	c.Node_number = configuration.Node_number
	c.Seed_node = configuration.Seed_node
	c.Service_specific_data = configuration.Service_specific_data
}

func (c *CoordActor) SendNewNodeResponse(m Message) {
	fmt.Println("Inside Response")
	var NewMessage Message
	fromNode := m.FromNode
	fmt.Println(c)
	for _, node := range c.Node_listeners {
		if node.Node_listner_url == fromNode {
			NewMessage.ContentData.Action = "New_Node_Response"
			NewMessage.FromNode = c.Node_addr
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

func (c *CoordActor) AddNewNodeMessageHandler() {
	var msgHandler MessageHandler
	msgHandler.MessagePacket.ContentData.Action = "New_Node"
	msgHandler.ServiceHandler = c.SendNewNodeResponse
	c.Cluster_op_msg_handler.AddMessageHandler(msgHandler)
}
