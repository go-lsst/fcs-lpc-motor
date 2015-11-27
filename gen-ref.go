// +build ignore

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	VM_NEGATIVE_REF_CLAMP1    = 50000.00 // Hz
	VM_POSITIVE_REF_CLAMP1    = 50000.00 // Hz
	VM_ACCEL_RATE             = 3200.000
	VM_MOTOR1_CURRENT_LIMIT   = 1000.0
	VM_SPEED                  = 50000.0
	VM_SPEED_FREQ_REF         = 50000.0
	VM_DRIVE_CURRENT_UNIPOLAR = 99999.999
	VM_DRIVE_CURRENT          = 99999.999
	VM_AC_VOLTAGE_SET         = 690
	VM_AC_VOLTAGE             = 930
	VM_RATED_CURRENT          = 99999.999
)

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

type Params []Param

func (p Params) Len() int           { return len(p) }
func (p Params) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Params) Less(i, j int) bool { return p[i].mbreg() < p[j].mbreg() }

func main() {
	flag.Parse()
	log.SetPrefix("[gen-ref] ")
	log.SetFlags(0)

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.Comment = '#'

	var params []Param
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		p := parseRecord(record)
		log.Printf("%#v\n", p)
		params = append(params, p)
	}

	if len(params) == 0 {
		log.Fatalf("no parameters!\n")
	}

	sort.Sort(Params(params))

	o, err := os.Create(flag.Arg(1))
	if err != nil {
		log.Fatalf("could not create file: %v\n", err)
	}
	defer o.Close()

	fmt.Fprintf(o, "package main\n\nfunc init() {\n")
	fmt.Fprintf(o, "\tparams = []Param{\n")

	for _, p := range params {
		fmt.Fprintf(
			o,
			"\t\t{menu: %d, index: %d, title: %q, defval: %q, rw: %v, size: %d},\n",
			p.menu, p.index, p.title, p.defval, p.rw, p.size,
		)
	}
	fmt.Fprintf(o, "\t}\n}\n")
	err = o.Close()
	if err != nil {
		log.Fatalf("error closing output file: %v\n", err)
	}
}

func parseRecord(data []string) Param {
	toks := strings.Split(data[0], ".")
	menu, err := strconv.Atoi(toks[0])
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	index, err := strconv.Atoi(toks[1])
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	rw := strings.TrimSpace(data[4]) == "rw"
	size, err := strconv.Atoi(strings.TrimSpace(data[3]))
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	return Param{
		menu:   menu,
		index:  index,
		title:  strings.TrimSpace(data[1]),
		defval: strings.TrimSpace(data[2]),
		size:   size,
		rw:     rw,
	}
}
