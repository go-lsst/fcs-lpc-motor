package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	codec  = binary.BigEndian
	params []Param
)

func main() {
	flag.Parse()

	if flag.NArg() == 1 {
		err := dispatch(flag.Arg(0))
		if err != nil {
			log.Fatalf("error dispatching [%s]: %v\n", flag.Arg(0), err)
		}
		log.Printf("\n")
		log.Printf("bye.\n")
		os.Exit(0)
	}

	params := []Parameter{
		{Menu: 0, Index: 1, Size: 1},
		{Menu: 0, Index: 2, Size: 1},
		{0, 3, 1},
		{0, 4, 1},
		{0, 10, 1}, // motor RPM
		{0, 31, 1},
		{0, 37, 2},
		{0, 48, 1}, // drive-mode (open-loop/RFC-A/RFC-s/regen)
		{0, 50, 1}, // software version
		{0, 16, 1},
		{0, 19, 1},
		{0, 20, 1},
		{0, 21, 1},
		{0, 22, 1},
		{0, 23, 1},
		{0, 24, 1},
		{0, 25, 1},
		{0, 27, 1},
		{0, 28, 1},
		{1, 6, 1},
		{1, 9, 1},
		{1, 10, 1},
		{1, 11, 1},
		{1, 21, 1}, // selected speed reference ?

		{5, 1, 1}, // output frequence
		{5, 5, 1},
		{5, 7, 1},
		{5, 8, 1},  // rated speed
		{5, 14, 1}, // drive mode (open-loop/RFC-A/RFC-S)
		{5, 53, 1}, // rotor temperature

		{11, 29, 1}, // software version
		{11, 30, 1}, // user security code
		{11, 31, 1}, // drive mode

		{21, 15, 1}, // motor-2 active (off=0, on=1)

		{22, 10, 1}, // motor RPM
		{22, 11, 1}, // output frequency
		{22, 50, 1}, // software version

		{24, 1, 1}, // module id
		{24, 2, 1}, // s/w version
		{24, 3, 1}, // h/w version
		{24, 4, 1}, // serial # LS
		{24, 5, 1}, // serial # MS
		{24, 6, 1}, // status
		{24, 7, 1}, // reset
		{24, 8, 1}, // default
		{24, 9, 1},
		{24, 10, 1}, // active ip
		{24, 11, 1}, // date code
		{24, 54, 1}, // drive date code
	}

	for _, m := range []Motor{
		NewMotor("134.158.125.223:502"),
		//NewMotor("134.158.125.224:502"),
	} {
		log.Printf(" === motor [%s] ===\n", m.Address)

		for _, p := range params {
			v, err := m.read(p)
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

		log.Printf("\n")
		log.Printf("-- end --\n\n")
	}

	m := NewMotor("134.158.125.223:502")
	for _, i := range []uint16{
		506, 507, 508,
	} {
		o, err := m.c.ReadHoldingRegisters(i, 1)
		log.Printf("Pr-%v: o=%v\terr=%v", NewParameter(i), o, err)
	}

	testParams(m)
}

func init() {
	log.SetPrefix("[fcs-motor] ")
	log.SetFlags(0)
}

type Param struct {
	menu   int
	index  int
	title  string
	defval string
	rw     bool
	size   int
	data   [2]byte
}

func (p *Param) mbreg() uint16 {
	return uint16(p.menu*100 + p.index - 1)
}

func testParams(m Motor) {
	log.Printf("")
	log.Printf(" --- test basic parameters ---\n")
	for _, p := range params {
		o, err := m.c.ReadHoldingRegisters(p.mbreg(), 1)
		log.Printf(
			"Pr-%02d.%03d: o=%v\t%q\terr=%v\n",
			p.menu, p.index, o, p.title, err,
		)
	}
}

func dispatch(name string) error {
	var err error
	switch name {
	case "shell":
		return runModbusShell()
	default:
		return fmt.Errorf("[%s] is not a valid command", name)
	}
	return err
}

func runModbusShell() error {
	sh := NewShell()
	defer sh.Close()
	return sh.run()
}
