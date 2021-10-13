package Testing

import (
	"runtime/debug"
	"testing"

	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
	coord "github.com/abhilashbss/distributed_coordinator/src/coordinator"
)

func TestCoordOldNodeJoining(t *testing.T) {

	logger.InitLogger("/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/log/logs.txt")
	coord1 := coord.CoordActor{}
	coord1.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_1.conf"
	coord1.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	coord1.LoadCoordinator()
	go coord1.Listen()

	coord2 := coord.CoordActor{}
	coord2.Node_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/node_init_2.conf"
	coord2.Cluster_conf_path = "/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/conf/cluster_meta.conf"
	go coord2.Listen()
	coord2.LoadCoordinator()

	debug.PrintStack()
	// if !reflect.DeepEqual(hashring, hashring_expected) {
	// 	t.Errorf("got %s, wanted %s", string(hashring_json), string(hashring_expected_json))
	// }

}
