package core

import (
	"elib/utils"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func WalkDirectory(path, suffix string) ([]string, error) {
	var nfoList []string
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), suffix) {
			nfoList = append(nfoList, path)
		}
		return nil
	})
	return nfoList, err
}

func DecodeNfo(path string) (*utils.Movie, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result utils.Movie
	if err := xml.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
