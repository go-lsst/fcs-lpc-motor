package main

import (
	"github.com/go-lsst/ncs/drivers/m702"
)

func init() {
	params = []m702.Parameter{
		{Index: [2]int{1, 5}, Title: "Jog reference", DefVal: "0.0", RW: true},
		{Index: [2]int{1, 6}, Title: "Maximum Reference Clamp", DefVal: "3000.0 rpm", RW: true},
		{Index: [2]int{1, 7}, Title: "Minimum Reference Clamp", DefVal: "0 rpm", RW: true},
		{Index: [2]int{1, 10}, Title: "Bipolar reference enable", DefVal: "0", RW: true},
		{Index: [2]int{1, 14}, Title: "Reference Selector", DefVal: "3", RW: true},
		{Index: [2]int{1, 21}, Title: "Preset reference 1", DefVal: "0.0", RW: true},
		{Index: [2]int{1, 22}, Title: "Preset reference 2", DefVal: "0.0", RW: true},
		{Index: [2]int{2, 2}, Title: "Ramp enable", DefVal: "1", RW: true},
		{Index: [2]int{2, 4}, Title: "Ramp mode select", DefVal: "1", RW: true},
		{Index: [2]int{2, 11}, Title: "Acceleration Rate 1", DefVal: "0.200s", RW: true},
		{Index: [2]int{2, 21}, Title: "Deceleration Rate 1", DefVal: "0.200s", RW: true},
		{Index: [2]int{3, 2}, Title: "Speed feedback", DefVal: "0.000001/rad", RW: true},
		{Index: [2]int{3, 8}, Title: "Overspeed threshold", DefVal: "0.0", RW: true},
		{Index: [2]int{3, 10}, Title: "Speed controller proportional gain Kp1", DefVal: "0.011 s/rad", RW: true},
		{Index: [2]int{3, 11}, Title: "Speed controller integral gain Ki1", DefVal: "1.00 s^2/rad", RW: true},
		{Index: [2]int{3, 12}, Title: "Speed controller differential feedback", DefVal: "0.000001/rad", RW: true},
		{Index: [2]int{3, 25}, Title: "Position feedback phase angle", DefVal: "n/a", RW: true},
		{Index: [2]int{3, 29}, Title: "P1 position", DefVal: "n/a", RW: false},
		{Index: [2]int{3, 34}, Title: "P1 rotary lines per revolution", DefVal: "4096", RW: true},
		{Index: [2]int{4, 1}, Title: "Current magnitude", DefVal: "n/a", RW: false},
		{Index: [2]int{4, 2}, Title: "Torque producing current", DefVal: "n/a", RW: false},
		{Index: [2]int{4, 7}, Title: "Symmetrical Current Limit", DefVal: "175%", RW: true},
		{Index: [2]int{4, 11}, Title: "Torque mode selector", DefVal: "0", RW: true},
		{Index: [2]int{4, 12}, Title: "Current reference filter time constant", DefVal: "0.0 ms", RW: true},
		{Index: [2]int{4, 13}, Title: "Current controller Kp gain", DefVal: "150", RW: true},
		{Index: [2]int{4, 14}, Title: "Current controller Ki gain", DefVal: "2000", RW: true},
		{Index: [2]int{4, 15}, Title: "Motor thermal time constant 1", DefVal: "89.0 s", RW: true},
		{Index: [2]int{5, 7}, Title: "Rated current", DefVal: "value of 11.032", RW: true},
		{Index: [2]int{5, 9}, Title: "Rated voltage", DefVal: "460V", RW: true},
		{Index: [2]int{5, 11}, Title: "Number of motor poles", DefVal: "6 Poles (3)", RW: true},
		{Index: [2]int{5, 12}, Title: "Auto-tune", DefVal: "0", RW: true},
		{Index: [2]int{5, 18}, Title: "Maximum switching frequence", DefVal: "6 kHz (3)", RW: true},
		{Index: [2]int{6, 13}, Title: "Enable auxiliary key", DefVal: "0", RW: true},
		{Index: [2]int{10, 37}, Title: "Action on trip detection", DefVal: "0", RW: true},
		{Index: [2]int{11, 29}, Title: "Software version", DefVal: "n/a", RW: false},
		{Index: [2]int{11, 30}, Title: "User security code", DefVal: "0", RW: true},
		{Index: [2]int{11, 31}, Title: "Drive mode", DefVal: "RFC-S (3)", RW: true},
		{Index: [2]int{11, 32}, Title: "Maximum heavy duty rating", DefVal: "n/a", RW: true},
		{Index: [2]int{11, 33}, Title: "Drive rated voltage", DefVal: "n/a", RW: false},
		{Index: [2]int{11, 36}, Title: "NV media card data previously loaded", DefVal: "n/a", RW: false},
		{Index: [2]int{11, 42}, Title: "Parameter cloning", DefVal: "0", RW: true},
		{Index: [2]int{11, 44}, Title: "User security status", DefVal: "Menu 0 (0)", RW: true},
		{Index: [2]int{24, 10}, Title: "Active IP address", DefVal: "n/a", RW: false},
	}
}
