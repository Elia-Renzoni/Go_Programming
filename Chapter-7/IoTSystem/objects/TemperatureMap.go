/**
*	@author Elia Renzoni
*	@date 12/01/2024
*	@brief OOP in Golang
 */

package objects

type temperatureMap struct {
	device
	temperatureData map[string]float64
}

func (t *temperatureMap) initMap() {
	t.temperatureData = make(map[string]float64)
}

func (t *temperatureMap) addDevice(device device) {
	t.temperatureData[device.locationDevice] = t.measureTemperature()
}

func (t *temperatureMap) updateTemperature(device device, newtemp float64) {
	for k := range t.temperatureData {
		if k == device.locationDevice {
			t.temperatureData[device.locationDevice] = newtemp
		}
	}
}

func (t *temperatureMap) getTemperature() (temperatures map[string]float64) {
	temperatures = t.temperatureData
	return
}
