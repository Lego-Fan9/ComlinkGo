package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Lego-Fan9/ComlinkGo"
)

func main() {
	// Create a struct with your settings
	comlink, err := ComlinkGo.GetComlink(&ComlinkGo.ComlinkSettings{
		ComlinkURL: "http://localhost:3000",
		HMAC: ComlinkGo.HMACSettings{
			AccessKey: "EXAMPLE",
			SecretKey: "EXAMPLE",
		},
	})

	if err != nil {
		panic(fmt.Sprintf("failed to create comlink: %v", err))
	}

	// Create your request body
	requestBody := ComlinkGo.RequestBody{
		Payload: ComlinkGo.Payload{
			AllyCode: "813479227",
		},
		Enums: false,
	}

	// Make the request
	playerData, err := comlink.Player(requestBody)
	if err != nil {
		panic(fmt.Sprintf("failed to get player: %v", err))
	}

	// OPTIONAL: Save it to file
	file, err := os.Create("example_player.json")
	if err != nil {
		panic(fmt.Sprintf("Error creating file: %v", err))
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(playerData)
	if err != nil {
		panic(fmt.Sprintf("Error encoding JSON: %v", err))
	}
}
