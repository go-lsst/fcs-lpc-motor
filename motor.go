package main

import (
	"fmt"
	"strconv"
	"strings"

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

type Parameter struct {
	Menu  int
	Index int
	Size  uint16
}

// NewParameterFromMenu creates a parameter from a menu.index string.
func NewParameterFromMenu(menu string) (Parameter, error) {
	var err error
	var p Parameter

	toks := strings.Split(menu, ".")
	m, err := strconv.Atoi(toks[0])
	if err != nil {
		return p, err
	}
	i, err := strconv.Atoi(toks[1])
	if err != nil {
		return p, err
	}
	return Parameter{Menu: m, Index: i, Size: 1}, err
}

// NewParameter creates a parameter from its modbus address register.
func NewParameter(reg uint16) Parameter {
	return Parameter{
		Menu:  int(reg / 100),
		Index: int(reg%100) + 1,
		Size:  1,
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
	o, err := m.c.ReadHoldingRegisters(p.ToModbus(), p.Size)
	if err != nil {
		return nil, err
	}
	return o, err
}

func (m *Motor) write(p Parameter, v []byte) ([]byte, error) {
	return m.c.WriteMultipleRegisters(p.ToModbus(), 1, v)
}
