package main

import (
	"os"

	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/hw"

	jww "github.com/spf13/jwalterweatherman"
)

func main() {
	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)

	adapters, err := hw.GetAdapters()
	if err != nil {
		jww.CRITICAL.Printf("Could not determine list of adapters: %v", err)
		os.Exit(1)
	}
	for _, adapter := range adapters {
		jww.INFO.Printf("Adapter found: %s", adapter.AdapterID)
	}

	ainfo, err := hw.GetAdapter("hci0")
	if err != nil {
		jww.CRITICAL.Printf("Problem instantiating adapter: %v", err)
		os.Exit(1)
	}

	jww.INFO.Printf("New BT management from adapter %s.", ainfo.AdapterID)
	btmgmt := hw.NewBtMgmt(ainfo.AdapterID)
	jww.INFO.Printf("Binpath: %s", btmgmt.BinPath)

	jww.INFO.Printf("Set Power off.")
	if err = btmgmt.SetPowered(false); err != nil {
		jww.CRITICAL.Printf("Could not power off BT mgmgt: %v", err)
		os.Exit(1)
	}
	jww.INFO.Printf("Set BT LE mode.")
	if err = btmgmt.SetLe(true); err != nil {
		jww.CRITICAL.Printf("Could not set BT LE mode: %v", err)
		os.Exit(1)
	}
	jww.INFO.Printf("Turn off BR/EDR support..")
	if err = btmgmt.SetBredr(false); err != nil {
		jww.CRITICAL.Printf("Could not turn off BR/EDR support: %v", err)
		os.Exit(1)
	}
	jww.INFO.Printf("Set Power on.")
	if err = btmgmt.SetPowered(true); err != nil {
		jww.CRITICAL.Printf("Could not power on BT mgmgt: %v", err)
		os.Exit(1)
	}
	jww.INFO.Printf("Instantiate adapter from adapter ID.")
	/* adapt */ _, err = adapter.NewAdapter1FromAdapterID(ainfo.AdapterID)
	if err != nil {
		jww.CRITICAL.Printf("Could not instantiate adapter from adapter ID '%s': %v", ainfo.AdapterID, err)
		os.Exit(1)
	}
}
