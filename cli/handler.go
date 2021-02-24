package cli

import (
	"flag"
	"fmt"
	"regexp"
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
	"first":  "",
	"second": "",
	"third":  "",
	"fourth": false,
}

func NewCommandHandler() *Command {
	return &Command{
		Args: CommandArgs,
	}
}

func (c *Command) HandleInput(args []string) error {

	var inputs []string

	for _, arg := range args {
		if !strings.Contains(arg, "-") {
			inputs = append(inputs, arg)
		}
	}

	if len(inputs) < 2 {
		return fmt.Errorf(`invalid parameter. example: analytics <last_minutes>m <path> <file_name>`)
	}

	argsWithoutProg := inputs[1:]

	if len(argsWithoutProg) < 2 {
		return fmt.Errorf(`invalid parameter. example: analytics <last_minutes>m <path>`)
	}

	tMinArg := argsWithoutProg[0]

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return fmt.Errorf(`invalid parameter %v`, err)
	}
	processedTminArg := reg.ReplaceAllString(tMinArg, "")

	tMin, err := strconv.ParseFloat(processedTminArg, 64)
	if err != nil {
		return fmt.Errorf(`invalid parameter minutes %v`, err)
	}

	c.Args["first"] = tMin
	c.Args["second"] = argsWithoutProg[1]
	c.Args["third"] = "http"
	c.Args["fourth"] = false

	if len(argsWithoutProg) >= 3 {
		c.Args["third"] = argsWithoutProg[2]
	}

	c.FlagHandle()

	return nil
}

func (c *Command) FlagHandle() error {
	boolPtr := flag.Bool("verbose", false, "verbose mode")
	//threadPtr := flag.Bool("thread", false, "thread mode")

	minutes := flag.String("t", "", "last n minutes")
	dirPtr := flag.String("d", "", "directory")
	filePrefix := flag.String("f", "", "file prefix")

	flag.Parse()

	if *boolPtr {
		c.Args["fourth"] = true
	}

	//if *threadPtr {
	//	c.Args["fourth"] = true
	//}

	if minutes != nil {
		reg, err := regexp.Compile("[^0-9]+")
		if err != nil {
			return fmt.Errorf(`invalid parameter %v`, err)
		}
		processedTminArg := reg.ReplaceAllString(*minutes, "")

		tMin, err := strconv.ParseFloat(processedTminArg, 64)
		if err != nil {
			return fmt.Errorf(`invalid parameter minutes %v`, err)
		}
		c.Args["first"] = tMin
	}

	if dirPtr != nil {
		if len(*dirPtr) < 1 {
			return fmt.Errorf(`invalid parameter dir. cant be empty`)
		}
		c.Args["second"] = *dirPtr
	}

	c.Args["third"] = "http"

	if filePrefix != nil {
		if len(*filePrefix) > 1 {
			c.Args["third"] = *filePrefix
		}
	}

	return nil
}
