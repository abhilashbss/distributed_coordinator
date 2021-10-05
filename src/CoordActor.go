package coordinator

type Coordinator struct {
	Node_count            int                `json:"Node_count"`
	Seed_node             int                `json:"Seed_node"`
	Node_listeners        []Node_url_mapping `json:"Node_url_mapping"`
	Node_number           int                `json:"Node_number"`
	Service_specific_data string             `json:"Service_specific_data"`
	MsgSender             MessageSender
}

func (c *Coordinator) RequestHandler(m Message) {

}

func (c *Coordinator) LoadCoordinator() {
	filepath := "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	var configuration Configuration
	configuration := LoadConf(filepath)

	if configuration.Node_number == configuration.Seed_node {
		c.Node_count = configuration.Node_count
		c.Node_listeners = configuration.Node_listeners
		c.Node_number = configuration.Node_number
		c.Seed_node = configuration.Seed_node
		c.Service_specific_data = configuration.Service_specific_data
	} else {
		c.Seed_node = configuration.Seed_node
		c.Node_listeners = configuration.Node_listeners
		c.RecieveConfFromSeed()
	}
}

//Resigter Actions for each service as per MessageHandlerMapping

func (c *Coordinator) LoadConfMessage() Message {
	var message Message
	message.FromNode = -1
	message.ToNode = c.Seed_node
	message.Content = `{"ServiceName" : "Coordinator", "Action" : "LoadConfig"}`
}

func (c *Coordinator) RecieveConfFromSeed() {
	c.MsgSender.SendMessage()
}
