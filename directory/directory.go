package directory

import (
	"fmt"
	"io/ioutil"
	"sort"

	log "github.com/sirupsen/logrus"
)

type (
	DirectoryInstance struct {
		FilePaths []string
	}

	DirectoryAction interface {
		Walk(string) error
		GetFilePaths() []string
	}
)

func NewDirectoryAction() (DirectoryAction, error) {
	return &DirectoryInstance{
		FilePaths: nil,
	}, nil
}

//scan directory for files//
func (di *DirectoryInstance) Walk(dirPath string) error {

	files, err := ioutil.ReadDir(dirPath)
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().After(files[j].ModTime())
	})

	for _, _file := range files {
		di.FilePaths = append(di.FilePaths, fmt.Sprintf(`%s/%s`, dirPath, _file.Name()))
	}

	if err != nil {
		log.Errorf(`error get directory _file list`)
		return err
	}
	for _, _file := range di.FilePaths {
		log.Infof(`found directory file %v`, _file)
	}

	return nil
}

func (di *DirectoryInstance) GetFilePaths() []string {
	return di.FilePaths
}
