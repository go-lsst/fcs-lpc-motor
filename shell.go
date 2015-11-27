package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/peterh/liner"
)

type Shell struct {
	shell  *liner.State
	prompt string
	cmds   map[string]shellCmd
	hist   string
	motor  Motor
}

func NewShell() *Shell {
	sh := &Shell{
		shell:  liner.NewLiner(),
		prompt: "mbus> ",
		hist:   filepath.Join(".", ".fcs_lpc_motor_history"),
		motor:  NewMotor("134.158.125.223:502"),
	}

	sh.shell.SetCtrlCAborts(true)
	sh.shell.SetCompleter(func(line string) (c []string) {
		for n := range sh.cmds {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	if f, err := os.Open(sh.hist); err == nil {
		sh.shell.ReadHistory(f)
		f.Close()
	}

	sh.cmds = map[string]shellCmd{
		"dump":  sh.cmdDump,
		"get":   sh.cmdGet,
		"motor": sh.cmdMotor,
		"quit":  sh.cmdQuit,
		"set":   sh.cmdSet,
	}
	return sh
}

type shellCmd func(args []string) error

func (sh *Shell) Close() error {
	if f, err := os.Create(sh.hist); err != nil {
		log.Print("error writing history file: ", err)
	} else {
		sh.shell.WriteHistory(f)
		f.Close()
	}
	fmt.Printf("\n")
	return sh.shell.Close()
}

func (sh *Shell) run() error {
	for {
		raw, err := sh.shell.Prompt(sh.prompt)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return err
		}
		// log.Printf("got: %q\n", raw)
		raw = strings.TrimSpace(raw)
		if raw == "" {
			continue
		}
		toks := strings.Split(raw, " ")
		err = sh.dispatch(toks)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return err
		}
		sh.shell.AppendHistory(raw)
	}

	return nil
}

func (sh *Shell) dispatch(toks []string) error {
	var err error
	fct, ok := sh.cmds[toks[0]]
	if !ok {
		err = fmt.Errorf("invalid command [%s]", toks[0])
		log.Printf("error: %v\n", err)
		return err
	}

	return fct(toks[1:])
}

func (sh *Shell) cmdQuit(args []string) error {
	return io.EOF
}

func (sh *Shell) cmdGet(args []string) error {
	param, err := sh.parseParam(args[0])
	if err != nil {
		return err
	}

	o, err := sh.motor.read(param)
	if err != nil {
		log.Printf("error: %v\n", err)
		err = nil
		return err
	}

	log.Printf(
		"Pr-%v: %s (%v)\n",
		param,
		displayBytes(o),
		codec.Uint16(o),
	)

	return err
}

func (sh *Shell) cmdSet(args []string) error {
	log.Printf(">>> %v\n", args)
	param, err := sh.parseParam(args[0])
	if err != nil {
		return err
	}
	vtype := "u16"
	if len(args) > 2 {
		vtype = args[2]
	}

	v := make([]byte, 2)

	switch vtype {
	case "u16", "uint16":
		vv, err := strconv.ParseUint(args[1], 10, 16)
		if err != nil {
			return err
		}
		codec.PutUint16(v, uint16(vv))

	default:
		return fmt.Errorf("cmd-set: invalid value-type (%v)", vtype)
	}

	log.Printf("set Pr-%v %s (%v)...\n", param, args[1], displayBytes(v))
	o, err := sh.motor.write(param, v)
	if err != nil {
		return err
	}
	log.Printf(
		"Pr-%v: %s (%v)\n",
		param,
		displayBytes(o),
		codec.Uint16(o),
	)

	return err
}

func (sh *Shell) cmdDump(args []string) error {
	var err error
	return err
}

func (sh *Shell) cmdMotor(args []string) error {
	switch len(args) {
	case 0:
		log.Printf("connected to [%s]\n", sh.motor.Address)
		return nil
	case 1:
		sh.motor = NewMotor(args[0])
		return nil
	default:
		return fmt.Errorf("cmd-motor: too many arguments (%d)", len(args))
	}
	return nil
}

func (sh *Shell) parseParam(arg string) (Parameter, error) {
	var err error
	var p Parameter

	if strings.Contains(arg, ".") {
		return NewParameterFromMenu(arg)
	}

	var reg uint64
	var base = 10
	if strings.HasPrefix(arg, "0x") {
		base = 16
		arg = arg[len("0x"):]
	}
	reg, err = strconv.ParseUint(arg, base, 64)
	if err != nil {
		return p, err
	}
	p = NewParameter(uint16(reg))
	return p, err
}

func displayBytes(o []byte) string {
	hex := make([]string, len(o))
	dec := make([]string, len(o))
	for i, v := range o {
		hex[i] = fmt.Sprintf("0x%02x", v)
		dec[i] = fmt.Sprintf("%3d", v)
	}

	return fmt.Sprintf("hex=%s dec=%s", hex, dec)
}
