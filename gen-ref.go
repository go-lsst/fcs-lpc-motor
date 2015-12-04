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

	"github.com/go-lsst/ncs/drivers/m702"
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

type Params []m702.Parameter

func (p Params) Len() int      { return len(p) }
func (p Params) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Params) Less(i, j int) bool {
	ii := int64(p[i].Index[0]*100000) + int64(p[i].MBReg())
	jj := int64(p[j].Index[0]*100000) + int64(p[j].MBReg())
	return ii < jj
}

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

	var params []m702.Parameter
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

	fmt.Fprintf(o, `package main

import (
	"github.com/go-lsst/ncs/drivers/m702"
)

func init() {
`)
	fmt.Fprintf(o, "\tparams = []m702.Parameter{\n")

	for _, p := range params {
		fmt.Fprintf(
			o,
			"\t\t{Index: [3]int{%d, %d, %d}, Title: %q, DefVal: %q, RW: %v},\n",
			p.Index[0], p.Index[1], p.Index[2], p.Title, p.DefVal, p.RW,
		)
	}
	fmt.Fprintf(o, "\t}\n}\n")
	err = o.Close()
	if err != nil {
		log.Fatalf("error closing output file: %v\n", err)
	}
}

func parseRecord(data []string) m702.Parameter {
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

	slot := 0 // FIXME
	return m702.Parameter{
		Index:  [3]int{slot, menu, index},
		Title:  strings.TrimSpace(data[1]),
		DefVal: strings.TrimSpace(data[2]),
		RW:     rw,
	}
}
