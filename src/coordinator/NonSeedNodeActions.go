package coordinator

import (
	"encoding/json"

	Util "github.com/abhilashbss/distributed_coordinator/src/util"
)

func (c *CoordActor) LoadNodeCoordinator() {
	filepath := "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init.conf"
	nodeConf, _ := Util.LoadNodeConf(filepath)
	c.Seed_addr = nodeConf.Seed_node
	c.Node_addr = nodeConf.Current_node
	c.SendConnectingMsgToSeedNode()
}

func (c *CoordActor) LoadConfMessage() Message {
	var message Message
	message.FromNode = c.Node_addr
	message.ToNode = c.Seed_addr
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

func (c *CoordActor) UpdateCoordMeta(m Message) {
	var msgCommunicator NewNodeCommunicator
	json.Unmarshal([]byte(m.ContentData.Data), &msgCommunicator)
	c.Node_count = msgCommunicator.Node_count
	c.Node_listeners = msgCommunicator.Node_listeners
	c.Service_specific_data = msgCommunicator.Service_specific_data
}
