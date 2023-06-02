//go:build ignore

package main

import (
	"encoding/json"
	"fmt"

	"go.melnyk.org/uuid"
	"gopkg.in/yaml.v3"
)

type data struct {
	Id uuid.UUID `json:"id"`
}

func main() {
	id := uuid.New()
	fmt.Printf("UUID value: %v\n", id)

	s := data{Id: id}

	output, err := json.Marshal(&s)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Json message: %s\n", string(output))

	var sj data
	err = json.Unmarshal(output, &sj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	output, err = yaml.Marshal(&s)
	fmt.Printf("Yaml message: %s\n", string(output))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	var sy data
	err = yaml.Unmarshal(output, &sy)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Unmarshaled structures:\n - %v\n - %v\n", sj, sy)
}
