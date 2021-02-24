package cli

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type (
	Command struct {
		Args map[string]interface{}
	}
)

const log_file_name = `http`

var CommandArgs = map[string]interface{}{
	"first": "",
	"second": "",
	"third": "",
	"fourth": "",
}

func NewCommandHandler() *Command {
	return &Command{
		Args: CommandArgs,
	}
}

func (c *Command) HandleInput (args []string) error {

	var inputs []string

	for _, arg := range args {
		if !strings.Contains(arg,"-") {
			inputs = append(inputs,arg)
		}
	}

	if len(inputs) < 2 {
		return fmt.Errorf(`invalid parameter. example: analytics <last_minutes>min <path> <file_name>`)
	}

	argsWithoutProg := inputs[1:]

	if len(argsWithoutProg) < 2 {
		return fmt.Errorf(`invalid parameter. example: analytics <last_minutes>min <path>`)
	}

	tMinArg := argsWithoutProg[0]
	tMin, err := strconv.ParseFloat(tMinArg,64)
	if err != nil {
		return fmt.Errorf(`error init directory action %v`, err)
	}

	c.Args["first"] = tMin
	c.Args["second"] = argsWithoutProg[1]
	c.Args["third"] =

	if len(argsWithoutProg) >= 3 {
		c.Args["third"] = argsWithoutProg[2]
	}

	FlagHandle()

	return nil
}

func FlagHandle() {
	boolPtr := flag.Bool("verbose", false, "verbose mode")
	flag.Parse()

	if *boolPtr {
		log.SetLevel(log.InfoLevel)
	}
}