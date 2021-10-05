package coordinator

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Node_count            int                `json:"Node_count"`
	Seed_node             int                `json:"Seed_node"`
	Node_listeners        []Node_url_mapping `json:"Node_url_mapping"`
	Node_number           int                `json:"Node_number"`
	Service_specific_data string             `json:"Service_specific_data"`
}

func LoadConf(filepath string) (Configuration, error) {

	file, _ := os.Open(filepath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
		return Configuration{}, err
	}
	fmt.Println(configuration)
	return configuration, nil
}
