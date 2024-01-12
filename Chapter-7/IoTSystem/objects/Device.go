/**
*	@author Elia Renzoni
*	@date 12/01/2023
*	@brief OOP in Golang
*
**/

package objects

import (
	"math/rand"
	"time"
)

type device struct {
	idDevice       string
	locationDevice string
}

func (d *device) measureTemperature() float64 {
	rand.Seed(time.Now().Unix())
	return rand.Float64()
}

func (d *device) newDevice(id, location string) *device {
	return &device{
		idDevice:       id,
		locationDevice: location,
	}
}
