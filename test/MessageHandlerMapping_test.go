package Testing

import (
	"fmt"
	"testing"

	coord "github.com/abhilashbss/distributed_coordinator/src/coordinator"
	messaging "github.com/abhilashbss/distributed_coordinator/src/messaging"
)

func testFunc(msg messaging.Message) {
	fmt.Println("Function call successful")
}

func TestMessageHandler(t *testing.T) {

	coord1 := coord.CoordActor{}
	message := messaging.Message{FromNode: "", ToNode: "", ServiceName: "", ContentData: messaging.Content{Action: "test", Data: ""}}

	var msgHandler messaging.MessageHandler
	msgHandler.MessagePacket.ContentData.Action = "test"
	msgHandler.ServiceHandler = testFunc
	coord1.Cluster_op_msg_handler.AddMessageHandler(msgHandler)

	coord1.ExecuteCoordinatorActionForMessage(message)

}
