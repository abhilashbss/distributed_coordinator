package coordinator

import (
	Util "github.com/abhilashbss/distributed_coordinator/src/util"
)

func (c *CoordActor) LoadCoordinator() {
	filepath := "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	configuration, _ := Util.LoadConf(filepath)

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
