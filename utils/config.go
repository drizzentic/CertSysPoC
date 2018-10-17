package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Configurations struct {
	Rpcusername string `yaml:"rpcusername"`
	Rpcpassword string `yaml:"rpcpassword"`
	Institution string `yaml:"institution"`
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

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	fmt.Print(dir)
}

func DeleteDirIfExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Remove(dir)
		if err != nil {
			panic(err)
		}
	}

	fmt.Print(dir)
}
