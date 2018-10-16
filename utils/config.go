package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Configurations struct {
	Rpcusername string `yaml:"rpcusername"`
	Rpcpassword string `yaml:"rpcpassword"`
}

//Get connection creds
func GetConnectionCredentials() Configurations {

	var c Configurations
	//Obtain values from yaml config
	yamlFile, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c

}
