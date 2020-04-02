package owner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	// ErrorReadFile happens on read file error.
	ErrorReadFile = errors.New("can't read file")
)

// GetALL get all owners.
func GetALL(filename string) ([]string, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Errorf("open file %s error: %w", filename, err)
		return nil, fmt.Errorf("%w:%w", ErrorReadFile, err)
	}
	result := strings.Split(string(bs), "\n")
	return result, nil
}
