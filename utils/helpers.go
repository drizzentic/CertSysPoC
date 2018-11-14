package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Configurations struct {
	Rpcusername      string `yaml:"rpcusername"`
	Rpcpassword      string `yaml:"rpcpassword"`
	Institution      string `yaml:"institution"`
	TransactionValue string `yaml:"transaction_value"`
}

//Get connection creds
func GetConfigs() Configurations {

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

func generateOPReturnHex() {

}
func GetHash(filename string) string {
	//Generate hash and create directory within the institution's folder

	h := sha256.New()
	h.Write([]byte(filename))
	a := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("filename", a)
	directoryHash := hex.EncodeToString(h.Sum(nil))

	return directoryHash
}

func GetDirectory(name string, admissionNumber string) (string,string) {
	fileFolder := GetHash(name + admissionNumber)

	filesDirectory := GetConfigs().Institution + "/" + fileFolder

	return filesDirectory,fileFolder
}

