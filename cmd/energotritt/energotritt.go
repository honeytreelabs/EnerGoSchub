package main

import (
	"encoding/hex"
	"flag"
	"os"
	"strings"

	jww "github.com/spf13/jwalterweatherman"
	"tinygo.org/x/bluetooth"
)

const (
	serviceSpec   = "e792361a-2c8a-4373-b365-e22c18dc1144"
	cableCharSpec = "ecbb79c5-6972-47ce-abbe-787a003c4f5d"
)

func main() {
	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)

	deviceMAC := flag.String("mac", "", "device MAC address")
	flag.Parse()

	if *deviceMAC == "" {
		jww.CRITICAL.Printf("The 'mac' argument is mandatory")
		os.Exit(1)
	}

	adapter := bluetooth.DefaultAdapter

	err := adapter.Enable()
	if err != nil {
		jww.CRITICAL.Printf("Could not enable BLE stack: %v", err)
		os.Exit(1)
	}

	mac, err := bluetooth.ParseMAC(*deviceMAC)
	if err != nil {
		jww.CRITICAL.Printf("Could not parse MAC '%s': %v", strings.ToUpper(*deviceMAC), err)
		os.Exit(1)
	}

	address := bluetooth.Address{MACAddress: bluetooth.MACAddress{MAC: mac}}
	address.SetRandom(false)
	device, err := adapter.Connect(address, bluetooth.ConnectionParams{})
	if err != nil {
		jww.CRITICAL.Printf("Failed to connect to '%s': %v", mac.String(), err)
		os.Exit(1)
	}

	jww.INFO.Printf("Connected to '%s'! Discovering services ...", mac.String())

	serviceUUID, err := bluetooth.ParseUUID(serviceSpec)
	if err != nil {
		jww.CRITICAL.Printf("Could not parse service specifier '%s': %v", serviceSpec, err)
		os.Exit(1)
	}
	services, err := device.DiscoverServices([]bluetooth.UUID{serviceUUID})
	if err != nil {
		jww.CRITICAL.Printf("Could not discover service with spec '%s': %v", serviceSpec, err)
		os.Exit(1)
	}
	service := services[0]

	jww.INFO.Printf("Found service: '%s'.", serviceSpec)

	cableCharUUID, err := bluetooth.ParseUUID(cableCharSpec)
	if err != nil {
		jww.CRITICAL.Printf("Failed to parse cable characteristic specifier '%s': %v", cableCharSpec, err)
		os.Exit(1)
	}

	cableChars, err := service.DiscoverCharacteristics([]bluetooth.UUID{cableCharUUID})
	if err != nil {
		jww.CRITICAL.Printf("Failed to discover characteristic '%s': %v", cableCharSpec, err)
		os.Exit(1)
	}
	cableChar := cableChars[0]

	jww.INFO.Printf("Found characteristic: '%s'.", cableCharSpec)

	bytes, err := hex.DecodeString("0200420a2c0802180120022a2431396234656165382d333239662d343266352d396239642d656361366133333734353964120a0a0865363761653239391a0622040a020806")
	if err != nil {
		jww.CRITICAL.Printf("Failed to decode amperage setting command: %v", err)
		os.Exit(1)
	}

	_, err = cableChar.WriteWithoutResponse(bytes)
	if err != nil {
		jww.CRITICAL.Printf("Failed to send payload: %v", err)
	}
	jww.INFO.Printf("Payload sent.")
}
