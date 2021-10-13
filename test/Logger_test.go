package Testing

import (
	"testing"

	logger "github.com/abhilashbss/distributed_coordinator/src/Logger"
)

func TestLogger(t *testing.T) {

	logger.InitLogger("/home/abhilashbss/go/src/github.com/abhilashbss/distributed_coordinator/log/logs_testing.txt")
	logger.ErrorLogger.Println("Error testing")
	logger.InfoLogger.Println("Info testing")
	logger.WarningLogger.Println("Warning testing")
	// if !reflect.DeepEqual(hashring, hashring_expected) {
	// 	t.Errorf("got %s, wanted %s", string(hashring_json), string(hashring_expected_json))
	// }
}
