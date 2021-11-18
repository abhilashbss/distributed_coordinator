# Distributed Coordinator

This is to tool that implements for multi-node communication and node addition cluster by self identification to seed node.

CLI: All cli commands needs to be run from /cli


## Service Implementation

Custom services can be implemented in services directory implementing 
- specific Handlers and Actions
- Implementing the following interface

```
type Service interface {
	AddHandlers()
	AddNodeOperations(Node_addr string, Node_listeners []CommonConfig.Node_url_mapping, MsgSender messaging.MessageSender)
	GetServiceName() string
	GetServiceHandler() messaging.MessageHandlerGroup
}
```

- Initialize the Service with the following
```
exService2 := &service.ExampleService{}
	exService2.AddHandlers()
	coord2.Service_message_processor.AddService(exService2)
```

