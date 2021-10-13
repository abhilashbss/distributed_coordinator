package CommonConfig

type Node_url_mapping struct {
	Node_id          int    `json:"Node_id"`
	Node_listner_url string `json:"Node_listener_url"`
}

type SeedConfiguration struct {
	Node_count            int                `json:"Node_count"`
	Seed_node             int                `json:"Seed_node"`
	Node_listeners        []Node_url_mapping `json:"Node_listeners"`
	Node_number           int                `json:"Node_number"`
	Service_specific_data string             `json:"Service_specific_data"`
}

type NodeConfiguration struct {
	Seed_node    string `json:"Seed_node"`
	Current_node string `json:"Current_node"`
}
