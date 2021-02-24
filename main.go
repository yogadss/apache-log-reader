package main

import (
	"accelbyte.test/cli"
	"accelbyte.test/directory/file"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

const (
	name      = `/usr/local/var/log/httpd`
	chunkSize = 2
)

func init()  {
	log.SetLevel(log.ErrorLevel)
}

func main() {

	start := time.Now()

	cmd := cli.NewCommandHandler()

	err := cmd.HandleInput(os.Args)
	if err != nil {
		log.Errorf(`error init directory action %v`, err)
	}

	defer func() {
		elapsed := time.Since(start)

		log.Infof("search took %s", elapsed)
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

	var (
		wg sync.WaitGroup
		chunks [][]string
	)

	items := fileAction.GetFilePaths()
	//cErr := make(chan error)
	cBreak := make(chan bool)

	go func(cs chan bool) {
		for {
			select {
			case <-cs:
				log.Infof("end signal fired. exited properly")
				elapsed := time.Since(start)
				log.Infof("search took %s", elapsed)
				os.Exit(1)
				return
			}
		}
	}(cBreak)

	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	for _, chunkChilds := range chunks {
		//errors := make([]error, len(chunkChilds))

		for _, chunkItem := range chunkChilds {

			wg.Add(1)

			go func(path string) {
				err := fileAction.ParseByLineSpecificFileChan(path, "access_log", cmd.Args["first"].(float64), cBreak)
				if err != nil {
					log.Errorf(`%v`, err)
				}
				defer wg.Done()
			}(chunkItem)
		}

		wg.Wait()

	}
}
