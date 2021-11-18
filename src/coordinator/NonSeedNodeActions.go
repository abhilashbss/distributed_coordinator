package coordinator

import (
	"encoding/json"

	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
	messaging "github.com/abhilashbss/distributed_coordinator/src/messaging"
	Util "github.com/abhilashbss/distributed_coordinator/src/util"
)

func (c *CoordActor) LoadNodeCoordinator() {
	filepath := c.Node_conf_path
	nodeConf, _ := Util.LoadNodeConf(filepath)
	c.Seed_addr = nodeConf.Seed_node
	c.Node_addr = nodeConf.Current_node
}

func (c *CoordActor) LoadConfMessage() messaging.Message {
	var message messaging.Message
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

type NewNodeCommunicator struct {
	Node_count            int                             `json:"Node_count"`
	Node_listeners        []CommonConfig.Node_url_mapping `json:"Node_listeners"`
	Service_specific_data string                          `json:"Service_specific_data"`
}

func (c *CoordActor) UpdateCoordMeta(m messaging.Message) {
	var msgCommunicator NewNodeCommunicator
	json.Unmarshal([]byte(m.ContentData.Data), &msgCommunicator)
	c.Node_count = msgCommunicator.Node_count
	c.Node_listeners = msgCommunicator.Node_listeners
	c.Service_specific_data = msgCommunicator.Service_specific_data
}

func (c *CoordActor) AddNewNodeResponseMessageHandler() {
	var msgHandler messaging.MessageHandler
	msgHandler.MessagePacket.ContentData.Action = "New_Node_Response"
	msgHandler.ServiceHandler = c.UpdateCoordMeta
	c.Cluster_op_msg_handler.AddMessageHandler(msgHandler)
}
