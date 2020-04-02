package cmd

import (
	"os"
	"reflect"
)

var envPrefix = "INPUT_"

// Input of the action.
type Input struct {
	Times           string `env:"TIMES"`
	GithubAuthToken string `env:"GITHUB_AUTH_TOKEN"`
	Owners          string `env:"OWNERS"`
}

// LoadFromEnv load env to input.
func (in *Input) LoadFromEnv() {
	num := reflect.ValueOf(in).Elem().NumField()
	for i := 0; i < num; i++ {
		tField := reflect.TypeOf(in).Elem().Field(i)
		vField := reflect.ValueOf(in).Elem().Field(i)
		value, ok := os.LookupEnv(envPrefix + tField.Tag.Get("env"))
		if ok {
			vField.Set(reflect.ValueOf(value))
		}
	}
}
