package coordinator

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
	Util "github.com/abhilashbss/distributed_coordinator/src/util"
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
	Node_conf_path            string
	Cluster_conf_path         string
	MsgSender                 MessageSender
	Cluster_op_msg_handler    MessageHandlerList
	Service_message_processor MessageHandlerList
	Router                    *gin.Engine
}

func (c *CoordActor) LoadCoordinator() {
	fmt.Println("Inside Load Coord")
	filepath := c.Node_conf_path
	nodeConf, _ := Util.LoadNodeConf(filepath)

	if nodeConf.Current_node == nodeConf.Seed_node {
		c.Seed_addr = nodeConf.Seed_node
		c.Node_addr = nodeConf.Current_node
		c.LoadSeedCoordinator()
		c.AddNewNodeMessageHandler()
	} else {
		c.LoadNodeCoordinator()
		c.AddNewNodeResponseMessageHandler()
	}
}

func (c *CoordActor) ExecuteCoordinatorActionForMessage(m Message) {
	c.Cluster_op_msg_handler.ExecuteForAction(m.ContentData.Action, m)
}

func (c *CoordActor) ExecuteServiceActionForMessage(m Message) {
	c.Service_message_processor.ExecuteForAction(m.ContentData.Action, m)
}

func (c *CoordActor) Listen() {
	c.Router = gin.Default()
	c.Router.POST("/service_request", func(con *gin.Context) {
		fmt.Println(con)
		var message Message
		if err := con.BindJSON(&message); err != nil {
			return
		}
		con.BindJSON(&message)
		logger.InfoLogger.Println("Service Message recieved")
		logger.InfoLogger.Println(message)
		c.ExecuteServiceActionForMessage(message)
		con.JSON(http.StatusOK, gin.H{})
	})

	c.Router.POST("/coordinator_request", func(con *gin.Context) {
		var message Message
		if err := con.BindJSON(&message); err != nil {
			return
		}
		con.BindJSON(&message)
		logger.InfoLogger.Println("Coordinator Message recieved")
		logger.InfoLogger.Println(message)
		c.ExecuteCoordinatorActionForMessage(message)
		con.JSON(http.StatusOK, gin.H{})
	})

	c.Router.Run(c.Node_addr)
}
