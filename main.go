package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

func main() {
	secret := make(map[string]interface{})

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("error: %w", err)
	}

	err = yaml.Unmarshal(stdin, &secret)
	if err != nil {
		log.Fatal("error: %w", err)
	}

	stringData := make(map[string]string)
	iter := reflect.ValueOf(secret["data"]).MapRange()
	for iter.Next() {
		key := iter.Key().Interface().(string)
		value := iter.Value().Interface().(string)
		item, err := base64.StdEncoding.DecodeString(string(value))
		if err != nil {
			log.Fatal("error: %w", err)
		}
		stringData[key] = string(item)
	}

	delete(secret, "data")
	secret["stringData"] = stringData

	var buf bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&buf)
	yamlEncoder.SetIndent(2)
	err = yamlEncoder.Encode(&secret)
	if err != nil {
		log.Fatal("error: %w", err)
	}

	fmt.Print(buf.String())
}
