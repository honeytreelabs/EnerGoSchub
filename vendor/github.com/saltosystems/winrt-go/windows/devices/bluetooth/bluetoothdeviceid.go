// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package bluetooth

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureBluetoothDeviceId string = "rc(Windows.Devices.Bluetooth.BluetoothDeviceId;{c17949af-57c1-4642-bcce-e6c06b20ae76})"

type BluetoothDeviceId struct {
	ole.IUnknown
}

func (impl *BluetoothDeviceId) GetId() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothDeviceId))
	defer itf.Release()
	v := (*iBluetoothDeviceId)(unsafe.Pointer(itf))
	return v.GetId()
}

func (impl *BluetoothDeviceId) GetIsClassicDevice() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothDeviceId))
	defer itf.Release()
	v := (*iBluetoothDeviceId)(unsafe.Pointer(itf))
	return v.GetIsClassicDevice()
}

func (impl *BluetoothDeviceId) GetIsLowEnergyDevice() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothDeviceId))
	defer itf.Release()
	v := (*iBluetoothDeviceId)(unsafe.Pointer(itf))
	return v.GetIsLowEnergyDevice()
}

const GUIDiBluetoothDeviceId string = "c17949af-57c1-4642-bcce-e6c06b20ae76"
const SignatureiBluetoothDeviceId string = "{c17949af-57c1-4642-bcce-e6c06b20ae76}"

type iBluetoothDeviceId struct {
	ole.IInspectable
}

type iBluetoothDeviceIdVtbl struct {
	ole.IInspectableVtbl

	GetId                uintptr
	GetIsClassicDevice   uintptr
	GetIsLowEnergyDevice uintptr
}

func (v *iBluetoothDeviceId) VTable() *iBluetoothDeviceIdVtbl {
	return (*iBluetoothDeviceIdVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iBluetoothDeviceId) GetId() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetId,
		uintptr(unsafe.Pointer(v)),        // this
		uintptr(unsafe.Pointer(&outHStr)), // out string
	)

	if hr != 0 {
		return "", ole.NewError(hr)
	}

	out := outHStr.String()
	ole.DeleteHString(outHStr)
	return out, nil
}

func (v *iBluetoothDeviceId) GetIsClassicDevice() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsClassicDevice,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iBluetoothDeviceId) GetIsLowEnergyDevice() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsLowEnergyDevice,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}
