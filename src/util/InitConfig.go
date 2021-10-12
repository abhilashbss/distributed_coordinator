package util

import (
	"encoding/json"
	"fmt"
	"os"

	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
)

func LoadConf(filepath string) (CommonConfig.Configuration, error) {

	file, _ := os.Open(filepath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := CommonConfig.Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
		return CommonConfig.Configuration{}, err
	}
	fmt.Println(configuration)
	return configuration, nil
}
