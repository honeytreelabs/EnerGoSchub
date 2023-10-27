package main

import (
	"github.com/rivo/tview"
	"tinygo.org/x/bluetooth"
)

func main() {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	adapter := bluetooth.DefaultAdapter
	adapter.Enable()
	app := tview.NewApplication()
	grid := tview.NewGrid().SetBorders(true)
	macsView := newPrimitive("Menu")
	sideBar := newPrimitive("Side Bar")

	var macs map[string]struct{}
	list := tview.NewList()
	list.
		AddItem("Start scan", "Selecting this starts a scan", 's', func() {
			macs = make(map[string]struct{})
			textView := macsView.(*tview.TextView)
			textView.SetText("")
			go adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
				macs[device.Address.String()] = struct{}{}
			})
		}).
		AddItem("Stop scan", "Selecting this stops the scan", 't', func() {
			adapter.StopScan()
			text := ""
			textView := macsView.(*tview.TextView)
			for mac := range macs {
				text = mac + "\n" + text
			}
			textView.SetText(text)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(list /*     */, 0, 0, 1, 1, 0, 0, false).
		AddItem(macsView /*  */, 1, 0, 1, 1, 0, 0, false).
		AddItem(sideBar /*   */, 2, 0, 1, 1, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(list /*     */, 0, 0, 1, 1, 0, 150, false).
		AddItem(macsView /*  */, 0, 1, 1, 1, 0, 150, false).
		AddItem(sideBar /*   */, 0, 2, 1, 1, 0, 150, false)

	if err := app.SetRoot(grid, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
