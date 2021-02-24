package main

import (
	"fmt"
	"log-search/cli"
	"log-search/directory/file"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	name      = `/usr/local/var/log/httpd`
	chunkSize = 2
)

func init() {
	log.SetLevel(log.FatalLevel)
}

func main() {

	start := time.Now()

	cmd := cli.NewCommandHandler()

	err := cmd.FlagHandle()
	if err != nil {
		log.Fatalf(`error input %v`, err)
	}

	if cmd.Args["fourth"].(bool) {
		log.SetLevel(log.InfoLevel)
	}

	//print elapsed time//
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("search took %s \n", elapsed)
	}()

	fileAction, err := file.NewFileAction()
	if err != nil {
		log.Fatalf(`error init directory action %v`, err)
	}

	err = fileAction.Walk(cmd.Args["second"].(string))
	if err != nil {
		log.Errorf(`error get list directory action %v`, err)
		os.Exit(1)
	}

	log.Println(`======================================`)

	items := fileAction.GetFilePaths()
	//cErr := make(chan error)
	cBreak := make(chan bool)

	//func to be executed when break signal is fired//
	go func(cs chan bool) {
		for {
			select {
			case <-cs:
				fmt.Printf("end signal fired. exited properly \n")
				elapsed := time.Since(start)
				fmt.Printf("search took %s \n", elapsed)
				os.Exit(1)
			}
		}
	}(cBreak)

	sequenceProc(fileAction, cmd, cBreak, items)
}

func sequenceProc(fileAction *file.FileInstance, cmd *cli.Command, cBreak chan bool, items []string) {
	for _, item := range items {
		err := fileAction.ParseByLineSpecificFileChan(item, cmd.Args["third"].(string), cmd.Args["first"].(float64), cBreak)
		if err != nil {
			log.Warnf(`%v`, err)
		}
	}
}
