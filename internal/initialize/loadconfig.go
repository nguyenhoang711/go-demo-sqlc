package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/hoangnguyen/demo-sqlc/global"
	"gopkg.in/yaml.v2"
)

func LoadConfig() {
	// yamlFile, err := os.ReadFile(fmt.Sprintf("./configs/%s.yaml", env))
	yamlFile, err := os.ReadFile("./configs/local.yaml")
	if err != nil {
        log.Fatalf("Error reading YAML file: %v", err)
    }

	// Parse the YAML file into the Config structure
	err = yaml.Unmarshal(yamlFile, &global.Config)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %v", err)
	}
	// read server configuration
	fmt.Println("Server port ::", global.Config.Server.Port)
}