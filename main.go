package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	var root yaml.Node
	fields := make(map[string]yaml.Node)

	yamlDecoder := yaml.NewDecoder(os.Stdin)
	err := yamlDecoder.Decode(&root)
	if err != nil {
		log.Fatal(err)
	}

	err = root.Decode(&fields)
	if err != nil {
		log.Fatal("error: %w", err)
	}

	data := make(map[string]string)
	dataNode := fields["data"]
	err = dataNode.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	stringData := make(map[string]string)
	stringDataNode := fields["stringData"]
	err = stringDataNode.Decode(&stringData)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range data {
		val, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			log.Fatal(err)
		}
		stringData[k] = string(val)
	}

	err = stringDataNode.Encode(&stringData)
	if err != nil {
		log.Fatal(err)
	}
	fields["stringData"] = stringDataNode
	delete(fields, "data")

	err = root.Encode(&fields)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&buf)
	yamlEncoder.SetIndent(2)
	err = yamlEncoder.Encode(&root)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf.String())
}
