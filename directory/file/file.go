package file

import (
	"accelbyte.test/directory"
	"accelbyte.test/utils"
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

type (
	FileInstance struct {
		directory.DirectoryInstance
		Lines []string
	}
)

const regexFilePrefix = `^(%s\.\d+)$|^(%s)$|^(%s-\d+\.log)$`

func NewFileAction () (*FileInstance, error) {
	return &FileInstance{
		DirectoryInstance : directory.DirectoryInstance{FilePaths: nil},
		Lines: nil,
	},nil
}

func (fi *FileInstance) ParseByLine(filePath string) error {

	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("failed opening directory: %s", err)
		return err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Errorf("failed closing directory: %s", err)
		}
	}()

	txtlines,err := LineParser(file)
	if err != nil {
		return fmt.Errorf("failed parse line: %s", err)
	}

	fi.Lines = txtlines

	return nil
}

func (fi *FileInstance) ParseByLineSpecificFile(filePath string, filePrefix string, lastMinutes float64) error {

	isTargetFile, err := CheckIfTargetFile(filePath,filePrefix)

	if !isTargetFile {
		return fmt.Errorf(`%s is not a target file`, filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("failed opening directory: %s", err)
		return err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Errorf("failed closing directory: %s", err)
		}
	}()

	txtlines,err := LineParser(file)
	if err != nil {
		return fmt.Errorf("failed parse line: %s", err)
	}

	for _, eachline := range txtlines {
		diffTime, err := utils.LogTimeDiff(eachline)
		if err != nil {
			log.Errorf(`failed to check log time diff %v skipping...`, err)
			continue
		}

		if diffTime <= lastMinutes {
			fmt.Println(eachline)
		}
	}

	fi.Lines = txtlines

	return nil
}

func (fi *FileInstance) ParseByLineSpecificFileChan(filePath string, filePrefix string, lastMinutes float64 ,bS chan bool) error {

	isTargetFile, err := CheckIfTargetFile(filePath,filePrefix)

	if !isTargetFile {
		return fmt.Errorf(`%s is not a target file`, filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed opening directory: %s", err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Errorf("failed closing directory: %s", err)
		}
	}()

	txtlines,err := LineParser(file)
	if err != nil {
		return fmt.Errorf("failed parse line: %s", err)
	}

	for _, eachline := range txtlines {
		diffTime, err := utils.LogTimeDiff(eachline)
		if err != nil {
			log.Errorf(`failed to check log time diff %v skipping...`, err)
			continue
		}

		if diffTime > lastMinutes {
			bS <- true
		}

		if diffTime <= lastMinutes {
			fmt.Println(eachline)
		}
	}

	log.Infof(`processed file %s ....`, filePath)

	return nil
}

func LineParser(file io.Reader) ([]string,error) {

	var txtlines []string
	
	if file == nil {
		return nil, fmt.Errorf(`invalid or empty file`)
	}
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	for _, eachline := range txtlines {
		log.Printf(eachline)
	}

	return txtlines,nil
}

func CheckIfTargetFile(filePath string,filePrefix string) (bool,error) {
	var isTarget bool

	if len(filePath) < 1 {
		return false, fmt.Errorf(`filePath length is zero`)
	}

	if len(filePrefix) < 1 {
		return false, fmt.Errorf(`filePrefix length is zero`)
	}

	fileName := filepath.Base(filePath)

	regexPattern := fmt.Sprintf(regexFilePrefix,filePrefix,filePrefix,filePrefix)

	re1, err := regexp.Compile(regexPattern) // error if regexp invalid
	if err != nil {
		log.Errorf("failed to compile regex 1: %s", err)
		return isTarget,err
	}

	isTargetFile := re1.MatchString(fileName)

	return isTargetFile,nil
}
