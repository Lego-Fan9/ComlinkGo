package tests

import (
	"testing"

	"github.com/Lego-Fan9/ComlinkGo"
)

func TestGameData(t *testing.T) {
	settings := &ComlinkGo.ComlinkSettings{ComlinkURL: *ComlinkURL}
	comlink, err := ComlinkGo.GetComlink(settings)
	if err != nil {
		t.Error(err)
	}

	metadata, err := comlink.Metadata(ComlinkGo.RequestBody{})
	if err != nil {
		t.Error(err)
	}

	_, ok := metadata["latestGamedataVersion"]
	if !ok {
		t.Error("latestGamedataVersion version didn't exist")
		return
	}

	version, ok := metadata["latestGamedataVersion"].(string)
	if !ok {
		t.Error("latestGamedataVersion exists but is not a string")
		return
	}

	payload := ComlinkGo.RequestBody{
		Payload: &ComlinkGo.Payload{
			Version: &version,
		},
	}

	_, err = comlink.GameData(payload)
	if err != nil {
		t.Error(err)
	}
}
