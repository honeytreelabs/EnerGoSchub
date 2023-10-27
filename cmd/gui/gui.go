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
	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(50, 0, 50).
		SetBorders(true).
		AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	menu := newPrimitive("Menu")
	sideBar := newPrimitive("Side Bar")

	var macs map[string]struct{}
	list := tview.NewList()
	list.
		AddItem("Start scan", "Selecting this starts a scan", 's', func() {
			macs = make(map[string]struct{})
			textView := menu.(*tview.TextView)
			textView.SetText("")
			go adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
				macs[device.Address.String()] = struct{}{}
			})
		}).
		AddItem("Stop scan", "Selecting this stops the scan", 't', func() {
			adapter.StopScan()
			text := ""
			textView := menu.(*tview.TextView)
			for mac := range macs {
				text = mac + "\n" + text
			}
			textView.SetText(text)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(list, 0, 0, 0, 0, 0, 0, false).
		AddItem(menu, 1, 0, 1, 3, 0, 0, false).
		AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(list, 1, 0, 1, 1, 0, 100, false).
		AddItem(menu, 1, 1, 1, 1, 0, 100, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	if err := app.SetRoot(grid, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
