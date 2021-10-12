package util

import (
	"encoding/json"
	"fmt"
	"os"

	CommonConfig "github.com/abhilashbss/distributed_coordinator/src/CommonConfig"
)

func LoadSeedConf(filepath string) (CommonConfig.SeedConfiguration, error) {

	file, _ := os.Open(filepath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := CommonConfig.SeedConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
		return CommonConfig.SeedConfiguration{}, err
	}
	fmt.Println(configuration)
	return configuration, nil
}

func LoadNodeConf(filepath string) (CommonConfig.NodeConfiguration, error) {

	file, _ := os.Open(filepath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := CommonConfig.NodeConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
		return CommonConfig.NodeConfiguration{}, err
	}
	fmt.Println(configuration)
	return configuration, nil
}
