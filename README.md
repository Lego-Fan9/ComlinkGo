# ComlinkGo

ComlinkGo is a wrapper for [swgoh-comlink](https://github.com/swgoh-utils/swgoh-comlink)

## Install
```bash
go get github.com/Lego-Fan9/ComlinkGo
```

## Usage
The following example will output the response of the /player endpoint:
```go
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
```
Please take note of the ComlinkURL and HMAC stuff near the top.

All comlink endpoints are avaliable. The ComlinkGo.RequestBody has fields for every possible input, just use it for all of them. The endpoints are all under nearly the same name as comlink has them, however the first letter is always capital. comlink.Player gets /player, comlink.GetEvents gets /getEvents, etc.

If you would like the raw *http.Response you can add a Raw to the end of the function call. Such as comlink.Player becomes comlink.PlayerRaw. If you use the raw functions, please note that I do not wrap the error. You get exactly what http.Do would give, unless it fails my retry logic (for more on retry logic please see httpclient/httpclient.go DoWithRetry())