package owner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	// ErrorReadFile happens on read file error.
	ErrorReadFile = errors.New("can't read file")
)

// GetALL get all owners.
func GetALL(filename string) ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadFile(filepath.Join(dir, filename))
	if err != nil {
		logrus.Errorf("open file %s error: %w", filename, err)
		return nil, fmt.Errorf("%w:%w", ErrorReadFile, err)
	}
	result := strings.Split(string(bs), "\n")
	return result, nil
}
