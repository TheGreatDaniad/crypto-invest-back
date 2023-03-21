package utils

import (
	"encoding/json"
	"fmt"
)

func PrintPretty(i interface{}) {
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
