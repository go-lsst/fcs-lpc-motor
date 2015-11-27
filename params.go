package main

func init() {
	params = []Param{
		{menu: 1, index: 5, title: "Jog reference", defval: "0.0", rw: true, size: 0},
		{menu: 1, index: 6, title: "Maximum Reference Clamp", defval: "3000.0 rpm", rw: true, size: 0},
		{menu: 1, index: 7, title: "Minimum Reference Clamp", defval: "0 rpm", rw: true, size: 0},
		{menu: 1, index: 10, title: "Bipolar reference enable", defval: "0", rw: true, size: 1},
		{menu: 1, index: 14, title: "Reference Selector", defval: "3", rw: true, size: 1},
		{menu: 1, index: 21, title: "Preset reference 1", defval: "0.0", rw: true, size: 0},
		{menu: 1, index: 22, title: "Preset reference 2", defval: "0.0", rw: true, size: 0},
		{menu: 2, index: 2, title: "Ramp enable", defval: "1", rw: true, size: 1},
		{menu: 2, index: 4, title: "Ramp mode select", defval: "1", rw: true, size: 0},
		{menu: 2, index: 11, title: "Acceleration Rate 1", defval: "0.200s", rw: true, size: 0},
		{menu: 2, index: 21, title: "Deceleration Rate 1", defval: "0.200s", rw: true, size: 0},
		{menu: 3, index: 2, title: "Speed feedback", defval: "0.000001/rad", rw: true, size: 0},
		{menu: 3, index: 8, title: "Overspeed threshold", defval: "0.0", rw: true, size: 0},
		{menu: 3, index: 10, title: "Speed controller proportional gain Kp1", defval: "0.011 s/rad", rw: true, size: 0},
		{menu: 3, index: 11, title: "Speed controller integral gain Ki1", defval: "1.00 s^2/rad", rw: true, size: 0},
		{menu: 3, index: 12, title: "Speed controller differential feedback", defval: "0.000001/rad", rw: true, size: 0},
		{menu: 3, index: 25, title: "Position feedback phase angle", defval: "n/a", rw: true, size: 0},
		{menu: 3, index: 29, title: "P1 position", defval: "n/a", rw: false, size: 0},
		{menu: 3, index: 34, title: "P1 rotary lines per revolution", defval: "4096", rw: true, size: 0},
		{menu: 4, index: 1, title: "Current magnitude", defval: "n/a", rw: false, size: 1},
		{menu: 4, index: 2, title: "Torque producing current", defval: "n/a", rw: false, size: 1},
		{menu: 4, index: 7, title: "Symmetrical Current Limit", defval: "175%", rw: true, size: 0},
		{menu: 4, index: 11, title: "Torque mode selector", defval: "0", rw: true, size: 0},
		{menu: 4, index: 12, title: "Current reference filter time constant", defval: "0.0 ms", rw: true, size: 1},
		{menu: 4, index: 13, title: "Current controller Kp gain", defval: "150", rw: true, size: 0},
		{menu: 4, index: 14, title: "Current controller Ki gain", defval: "2000", rw: true, size: 0},
		{menu: 4, index: 15, title: "Motor thermal time constant 1", defval: "89.0 s", rw: true, size: 0},
		{menu: 5, index: 7, title: "Rated current", defval: "value of 11.032", rw: true, size: 0},
		{menu: 5, index: 9, title: "Rated voltage", defval: "460V", rw: true, size: 0},
		{menu: 5, index: 11, title: "Number of motor poles", defval: "6 Poles (3)", rw: true, size: 0},
		{menu: 5, index: 12, title: "Auto-tune", defval: "0", rw: true, size: 0},
		{menu: 5, index: 18, title: "Maximum switching frequence", defval: "6 kHz (3)", rw: true, size: 0},
		{menu: 6, index: 13, title: "Enable auxiliary key", defval: "0", rw: true, size: 0},
		{menu: 10, index: 37, title: "Action on trip detection", defval: "0", rw: true, size: 0},
		{menu: 11, index: 29, title: "Software version", defval: "n/a", rw: false, size: 0},
		{menu: 11, index: 30, title: "User security code", defval: "0", rw: true, size: 4},
		{menu: 11, index: 31, title: "Drive mode", defval: "RFC-S (3)", rw: true, size: 0},
		{menu: 11, index: 32, title: "Maximum heavy duty rating", defval: "n/a", rw: true, size: 0},
		{menu: 11, index: 33, title: "Drive rated voltage", defval: "n/a", rw: false, size: 0},
		{menu: 11, index: 36, title: "NV media card data previously loaded", defval: "n/a", rw: false, size: 0},
		{menu: 11, index: 42, title: "Parameter cloning", defval: "0", rw: true, size: 0},
		{menu: 11, index: 44, title: "User security status", defval: "Menu 0 (0)", rw: true, size: 0},
		{menu: 24, index: 10, title: "Active IP address", defval: "n/a", rw: false, size: 4},
	}
}
