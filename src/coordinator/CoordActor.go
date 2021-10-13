package coordinator

import (
	"net/http"

	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
	Util "github.com/abhilashbss/distributed_coordinator/src/util"
	"github.com/gin-gonic/gin"
)

type CoordActor struct {
	Node_count                int                             `json:"Node_count"`
	Seed_node                 int                             `json:"Seed_node"`
	Node_listeners            []CommonConfig.Node_url_mapping `json:"Node_listeners"`
	Node_number               int                             `json:"Node_number"`
	Current_node_listener     string                          `json:"Current_node_listener"`
	Service_specific_data     string                          `json:"Service_specific_data"`
	Seed_addr                 string                          `json:"Seed_address"`
	Node_addr                 string                          `json:"Node_address"`
	MsgSender                 MessageSender
	Cluster_op_msg_handler    MessageHandlerList
	Service_message_processor MessageHandlerList
	Router                    *gin.Engine
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

func (c *CoordActor) ExecuteCoordinatorActionForMessage(m Message) {
	c.Cluster_op_msg_handler.ExecuteForAction(m.ContentData.Action, m)
}

func (c *CoordActor) ExecuteServiceActionForMessage(m Message) {
	c.Service_message_processor.ExecuteForAction(m.ContentData.Action, m)
}

func (c *CoordActor) Listen() {

	c.Router.Run(c.Node_addr)
	c.Router.POST("/service_request", func(con *gin.Context) {

		var message Message
		if err := con.BindJSON(&message); err != nil {
			return
		}
		con.BindJSON(&message)
		c.ExecuteServiceActionForMessage(message)
		con.JSON(http.StatusOK, gin.H{})
	})

	c.Router.POST("/coordinator_request", func(con *gin.Context) {

		var message Message
		if err := con.BindJSON(&message); err != nil {
			return
		}
		con.BindJSON(&message)
		c.ExecuteCoordinatorActionForMessage(message)
		con.JSON(http.StatusOK, gin.H{})
	})
}
