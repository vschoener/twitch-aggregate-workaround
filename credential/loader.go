package credential

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Loader interface
type Loader interface {
	// This method should load the source define as string to store element
	// in the interface
	Load(string, interface{}) error
}

// YAMLLoader is a YAML loader parameters
type YAMLLoader struct {
}

// Load will load the parameters from Yaml file and store it insie the
// storage
func (y YAMLLoader) Load(path string, storage interface{}) error {
	raw, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(raw, storage); err != nil {
		panic(err)
	}

	return err
}
