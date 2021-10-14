package coordinator

import messaging "github.com/abhilashbss/distributed_coordinator/src/messaging"

type ServiceMessageProcessor struct {
}

func (c *ServiceMessageProcessor) ExecuteActionForMessage(m messaging.Message) {
}
