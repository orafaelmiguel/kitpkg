package commands

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonFmtCommand struct{}

func (j JsonFmtCommand) Name() string {
	return "jsonfmt"
}

func (j JsonFmtCommand) Description() string {
	return "Formats and validates JSON"
}

func (j JsonFmtCommand) Execute(input string, args []string) string {
	var data []byte
	var err error

	if input != "" {
		data = []byte(input)
	} else if len(args) > 0 {
		data, err = os.ReadFile(args[0])
		if err != nil {
			return fmt.Sprintf("error reading file: %v", err)
		}
	} else {
		return "usage: jsonfmt <file> or pipe input"
	}

	var parsed interface{}

	err = json.Unmarshal(data, &parsed)
	if err != nil {
		return fmt.Sprintf("invalid JSON: %v", err)
	}

	formatted, err := json.MarshalIndent(parsed, "", "  ")
	if err != nil {
		return fmt.Sprintf("error formatting JSON: %v", err)
	}

	return string(formatted)
}