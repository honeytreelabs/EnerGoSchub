package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"tinygo.org/x/bluetooth"
)

func main() {
	sideBar := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Control")
	adapter := bluetooth.DefaultAdapter
	adapter.Enable()
	scanRunning := false
	app := tview.NewApplication()
	grid := tview.NewGrid().SetBorders(true)
	actionView := tview.NewList()
	scanResultsView := tview.NewTable().SetBorders(true).SetSelectable(true, false).SetSelectedFunc(func(row, column int) {
		app.SetFocus(actionView)
	})

	macs := NewMacStorage()
	actionView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTAB:
			app.SetFocus(scanResultsView)
			return nil
		}
		return event
	})
	scanResultsView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTAB:
			app.SetFocus(actionView)
			return nil
		}
		switch string(event.Rune()) {
		case "q":
			fallthrough
		case "Q":
			app.Stop()
			return nil
		}
		return event
	})

	startScan := func() {
		if scanRunning {
			return
		}
		scanRunning = true

		macs.Clear()
		go adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
			macs.Add(device)

			scanResultsView.Clear()
			for i, scanResult := range macs.GetAll(nil) {
				scanResultsView.SetCell(i, 0, tview.
					NewTableCell(fmt.Sprintf("%s @ %d", scanResult.Address.String(), scanResult.RSSI)).
					SetTextColor(tcell.ColorWhite).
					SetAlign(tview.AlignCenter).
					SetExpansion(1))
			}
			app.Draw()
		})
	}

	actionView.
		AddItem("Start scan", "Selecting this starts a scan", 's', startScan).
		AddItem("Stop scan", "Selecting this stops the scan", 't', func() {
			adapter.StopScan()
			scanRunning = false
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(actionView /*        */, 0, 0, 1, 1, 0, 0, false).
		AddItem(scanResultsView /*    */, 1, 0, 1, 1, 0, 0, false).
		AddItem(sideBar /*            */, 2, 0, 1, 1, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(actionView /*        */, 0, 0, 1, 1, 0, 150, false).
		AddItem(scanResultsView /*    */, 0, 1, 1, 1, 0, 150, false).
		AddItem(sideBar /*            */, 0, 2, 1, 1, 0, 150, false)

	app.SetRoot(grid, true).SetFocus(actionView)
	startScan()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
