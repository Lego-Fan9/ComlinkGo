package tests

import (
	"flag"
	"testing"

	"github.com/Lego-Fan9/ComlinkGo"
)

var ComlinkURL = flag.String("comlink", "", "A valid ComlinkURL for testing")

func TestEnums(t *testing.T) {
	settings := &ComlinkGo.ComlinkSettings{ComlinkURL: *ComlinkURL}
	comlink, err := ComlinkGo.GetComlink(settings)
	if err != nil {
		t.Error(err)
	}

	_, err = comlink.Enums()
	if err != nil {
		t.Error(err)
	}
}
