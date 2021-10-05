package coordinator

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Listeners struct {
	Cluster_op_msg_handler MessageHandlerList
	Service_msg_handler    MessageHandlerList
	Router                 *gin.Engine
	Coord_actor            coordinator
	Service_msg_service    ServiceMessagingService
}

// create getter and setters for all the objects

func (l *Listeners) Listen() {

	l.Router.POST("/service_request", func(c *gin.Context) {

		var message Message
		if err := c.BindJSON(&message); err != nil {
			return
		}
		c.BindJSON(&message)
		l.Service_msg_handler.RequestHandler(message)
		c.JSON(http.StatusOK)

	})

	l.Router.POST("/coordinator_request", func(c *gin.Context) {

		var message Message
		if err := c.BindJSON(&message); err != nil {
			return
		}
		c.BindJSON(&message)
		l.Coord_actor.RequestHandler(message)
		c.JSON(http.StatusOK)
	})
}

func (l *Listeners) StartListeners() {
	l.Router.Run("localhost:8080")
}
