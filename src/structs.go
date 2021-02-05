package main

import (
	"encoding/json"
	"fmt"
)

type courseMeta struct {
	Author string
	Level  string
	Rating float64
}

func main() {
	fmt.Println("structs")

	// basicStructInit()

	jsonDoc := `
        {
            "environment" : "asdf",
            "hostName" : ""
        }
	`
	conf := &configWithPointers{}
	json.Unmarshal([]byte(jsonDoc), conf)

	fmt.Println(jsonDoc)
	fmt.Println(conf)
}

// Using string pointers differs from empty strings in the sense that the value might not have been initialized at all.
type configWithPointers struct {
	Environment *string // pointer to string
	Version     *string
	HostName    *string
}

func (c *configWithPointers) String() string {
	envOut, verOut, hostOut := "<nil>", "<nil>", "<nil>"

	if c.Environment != nil { // Check for nil!
		envOut = *c.Environment
	}

	if c.Version != nil {
		verOut = *c.Version
	}

	if c.HostName != nil {
		hostOut = *c.HostName
	}

	return fmt.Sprintf("Environment: '%v'\nVersion:'%v'\nHostName: '%v'",
		envOut, verOut, hostOut)
}

func basicStructInit() {
	ThisCourse := courseMeta{
		Author: "Someone",
		Level:  "Aiai",
		Rating: 12.3,
	}

	fmt.Println(ThisCourse)
}
