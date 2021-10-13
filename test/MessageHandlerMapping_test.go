package Testing

import (
	"fmt"
	"testing"

	coord "github.com/abhilashbss/distributed_coordinator/src/coordinator"
)

func testFunc(msg coord.Message) {
	fmt.Println("Function call successful")
}

func TestMessageHandler(t *testing.T) {

	coord1 := coord.CoordActor{}
	message := coord.Message{FromNode: "", ToNode: "", ServiceName: "", ContentData: coord.Content{Action: "test", Data: ""}}

	var msgHandler coord.MessageHandler
	msgHandler.MessagePacket.ContentData.Action = "test"
	msgHandler.ServiceHandler = testFunc
	coord1.Cluster_op_msg_handler.AddMessageHandler(msgHandler)

	coord1.ExecuteCoordinatorActionForMessage(message)

}
