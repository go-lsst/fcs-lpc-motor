package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-lsst/ncs/drivers/m702"
)

var (
	codec  = binary.BigEndian
	params []m702.Parameter
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

	m := m702.New("134.158.125.223:502")
	testParams(m)
}

func init() {
	log.SetPrefix("[fcs-motor] ")
	log.SetFlags(0)
}

func testParams(m m702.Motor) {
	log.Printf("")
	log.Printf(" --- test basic parameters ---\n")
	for _, p := range params {
		err := m.ReadParam(&p)
		log.Printf(
			"Pr-%v: o=%v\t%q\terr=%v\n",
			p, p.Data[:], p.Title, err,
		)
	}
}

func dispatch(name string) error {
	var err error
	switch name {
	case "shell":
		return runModbusShell()
	case "web":
		return runWebServer()
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
