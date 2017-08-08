package main

import (
	// "fmt"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"log"
)

func main() {
	info := accessory.Info{
		Name:         "Yeelight Lightbulb",
		Manufacturer: "Yeelight",
	}

	acc := NewYeelight(info, "192.168.0.27", "55443")
	t, err := hc.NewIPTransport(hc.Config{Pin: "12341234"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
