package MSCreater

import (
	"fmt"
	"io/ioutil"
)

type Yaml struct {
	inPath, outPath string
}

func (y Yaml) Generate() error {
	yamlFile, err := ioutil.ReadFile(y.inPath)
	if err != nil {
		fmt.Printf("failed to read yaml file: %v\n", err)
		return err
	}
	var items Items

	return nil
}
