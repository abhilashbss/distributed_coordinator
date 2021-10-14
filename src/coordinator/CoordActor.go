package coordinator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
	messaging "github.com/abhilashbss/distributed_coordinator/src/messaging"
	service "github.com/abhilashbss/distributed_coordinator/src/service"
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
	MsgSender                 messaging.MessageSender
	Cluster_op_msg_handler    messaging.MessageHandlerGroup
	Service_message_processor service.ServiceGroup
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
	} else {
		c.LoadNodeCoordinator()
	}
}

func (c *CoordActor) ExecuteCoordinatorActionForMessage(m messaging.Message) {
	c.Cluster_op_msg_handler.ExecuteForAction(m.ContentData.Action, m)
}

func (c *CoordActor) ExecuteServiceActionForMessage(m messaging.Message) {
	c.Service_message_processor.ExecuteMessageAction(m)
}

func (c *CoordActor) Listen() {
	c.AddNewNodeMessageHandler()
	c.AddNewNodeResponseMessageHandler()

	c.Router = gin.Default()
	c.Router.POST("/service_request", func(con *gin.Context) {
		fmt.Println(con)
		var message messaging.Message
		if err := con.BindJSON(&message); err != nil {
			return
		}
		con.BindJSON(&message)
		messageJson, _ := json.Marshal(message)
		logger.InfoLogger.Println("Service Message recieved : " + string(messageJson))
		c.ExecuteServiceActionForMessage(message)
		con.JSON(http.StatusOK, gin.H{})
	})

	c.Router.POST("/coordinator_request", func(con *gin.Context) {
		var message messaging.Message
		if err := con.BindJSON(&message); err != nil {
			return
		}
		con.BindJSON(&message)
		messageJson, _ := json.Marshal(message)
		logger.InfoLogger.Println("Coordinator Message recieved : " + string(messageJson))
		c.ExecuteCoordinatorActionForMessage(message)
		logger.InfoLogger.Println("Current coordinator state : " + c.CurrentState())
		con.JSON(http.StatusOK, gin.H{})
	})

	c.Router.Run(c.Node_addr)
}

func (c *CoordActor) CurrentState() string {
	listenerJson, _ := json.Marshal(c.Node_listeners)
	state := "Node addr : " + c.Node_addr + " Nodes: " + strconv.Itoa(c.Node_count) + " Node_listeners : " + string(listenerJson)
	return string(state)
}
