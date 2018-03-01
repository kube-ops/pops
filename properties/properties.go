package properties

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// GetYAMLProperties retrieve properties from file
func GetYAMLProperties(path string, out interface{}) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, out)
	if err != nil {
		return errors.New("unmarshal " + path + ": yaml badly formatted")
	}
	return nil
}

// SetYAMLProperties write properties into file
func SetYAMLProperties(path string, in interface{}) error {
	file, err := yaml.Marshal(in)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, file, os.FileMode(uint32(0664)))
	return err
}
