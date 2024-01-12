/**
*	@author Elia Renzoni
*	@date 12/01/2024
*	@brief OOP in Golang
**/

package objects

import "fmt"

type iotSystem struct {
	devices
	temperatureMap
}

type devices struct {
	device
	nextDevice *devices
	head       *devices
}

func (o *iotSystem) registerNewDevice(id, location string) {
	var newDevice *devices = &devices{}
	newDevice.idDevice = id
	newDevice.locationDevice = location
	if o.head == newDevice {
		o.head = newDevice
	} else {
		lastNode := o.getLastDevice()
		lastNode.nextDevice = newDevice
	}
}

func (o *iotSystem) updateTemperatureData() {
	// TODO
}

func (o *iotSystem) updateTemperatures() {
	// TODO
}

func (o *iotSystem) getTemperature(location string) float64 {
	return o.temperatureData[location]
}

func (o *iotSystem) viewDevices() {
	for node := o.head; node != nil; node = node.nextDevice {
		fmt.Println("%s, %s", node.idDevice, node.locationDevice)
	}
}

func (o *iotSystem) getLastDevice() *devices {
	var node *devices
	for node := o.head; node != nil; node = node.nextDevice {
		node = node
	}
	return node
}
