package main

import (
	"sort"
	"strings"

	"tinygo.org/x/bluetooth"
)

type MacStorage struct {
	macs map[string]bluetooth.ScanResult
}

func NewMacStorage() *MacStorage {
	return &MacStorage{}
}

func (m *MacStorage) Clear() {
	m.macs = make(map[string]bluetooth.ScanResult)
}

func (m *MacStorage) Add(scanResult bluetooth.ScanResult) {
	m.macs[scanResult.Address.String()] = scanResult
}

func (m *MacStorage) GetAll(filter *string) []bluetooth.ScanResult {
	var result []bluetooth.ScanResult
	for mac := range m.macs {
		scanResult := m.macs[mac]
		if filter == nil || strings.HasPrefix(strings.ToLower(scanResult.Address.String()), strings.ToLower(*filter)) {
			result = append(result, m.macs[mac])
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Address.String() < result[j].Address.String()
	})
	return result
}
