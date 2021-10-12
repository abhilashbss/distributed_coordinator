package Testing

import (
	"fmt"
	"testing"

	coord "github.com/abhilashbss/distributed_coordinator/src/coordinator"
)

func printt(s coord.Message) {
	fmt.Printf(s.ContentData.Data + "printtttt")
}

//Print is working
func TestConfLoad(t *testing.T) {
	msg := coord.Message{ServiceName: "abs", FromNode: 1, ToNode: 2, ContentData: coord.Content{Action: "New_Node", Data: "fsd"}}

	m := coord.MessageHandler{}
	m.SetServiceHandler(printt)
	m.SetMessagePacket(msg)
	m.ExecuteMessageHandler()

	// if !reflect.DeepEqual(hashring, hashring_expected) {
	// 	t.Errorf("got %s, wanted %s", string(hashring_json), string(hashring_expected_json))
	// }
}
