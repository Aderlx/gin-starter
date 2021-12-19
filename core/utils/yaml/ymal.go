package yaml

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"
)

func ReadYaml(logConfig *interface{}, filePath string) interface{} {
	yamlFile, err := ioutil.ReadFile(filePath)
	log.Println("yamlFile:", yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	// var config = &log.LogConfig{}
	err = yaml.Unmarshal(yamlFile, logConfig)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	s := reflect.ValueOf(logConfig)
	fmt.Print(s)
	return logConfig

}
