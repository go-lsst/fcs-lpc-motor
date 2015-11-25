package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/goburrow/modbus"
)

/*
### register mode = standard
=> register address = (mm x 100) + ppp - 1
   where mm <= 162 && ppp <= 99

### register mode = modified
=> register address = (mm x 256) + ppp - 1
   where mm <= 63 && ppp <= 255
*/

var (
	codec = binary.BigEndian
)

type Parameter struct {
	Menu  int
	Index int
}

// NewParameter creates a parameter from its modbus address register.
func NewParameter(reg uint16) Parameter {
	return Parameter{
		Menu:  int(reg / 100),
		Index: int(reg%100) + 1,
	}
}

func (p Parameter) ToModbus() uint16 {
	return uint16(p.Menu*100 + p.Index - 1)
}

func (p Parameter) String() string {
	return fmt.Sprintf("%02d.%03d", p.Menu, p.Index)
}

type Motor struct {
	Address string
	c       modbus.Client
}

func NewMotor(addr string) Motor {
	return Motor{
		Address: addr,
		c:       modbus.TCPClient(addr),
	}
}

func (m *Motor) read(p Parameter) ([]byte, error) {
	o, err := m.c.ReadHoldingRegisters(p.ToModbus(), 1)
	return o, err
}

func main() {
	params := []Parameter{
		{Menu: 0, Index: 1},
		{Menu: 0, Index: 2},
		{0, 3},
		{0, 4},
		{0, 10}, // motor RPM
		{0, 31},
		{0, 37},
		{0, 48}, // drive-mode (open-loop/RFC-A/RFC-s/regen)
		{0, 50}, // software version
		{0, 16},
		{0, 19},
		{0, 20},
		{0, 21},
		{0, 22},
		{0, 23},
		{0, 24},
		{0, 25},
		{0, 27},
		{0, 28},
		{1, 6},
		{1, 9},
		{1, 10},
		{1, 11},
		{1, 21}, // selected speed reference ?

		{5, 1},  // output frequence
		{5, 4},  // motor RPM
		{5, 14}, // drive mode (open-loop/RFC-A/RFC-S)
		{5, 53}, // rotor temperature

		{11, 29}, // software version

		{21, 15}, // motor-2 active (off=0, on=1)

		{22, 10}, // motor RPM
		{22, 11}, // output frequency
		{22, 50}, // software version

	}

	m1 := NewMotor("134.158.125.223:502")

	for _, p := range params {
		v, err := m1.read(p)
		var vv uint64
		switch len(v) {
		case 2:
			vv = uint64(codec.Uint16(v))
		case 4:
			vv = uint64(codec.Uint32(v))
		case 8:
			vv = codec.Uint64(v)
		}
		if err != nil {
			log.Printf(
				"Pr-%s (%05d): err=%v\n",
				p, p.ToModbus(), err,
			)
		} else {
			log.Printf(
				"Pr-%s (%05d): %v %v\n",
				p, p.ToModbus(), v, vv,
			)
		}
	}
}

func init() {
	log.SetPrefix("[fcs-motor] ")
}
