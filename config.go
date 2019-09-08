package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

//Get is a proxy to C().Get
var Get func(key string) string

//Map represnets config
type Map map[string]string

//Get gets a key from config map
func (c Map) Get(key string) string {
	val, exists := (c)[key]
	if !exists {
		return ""
	}
	return val
}

//Set sets a key config
func (c *Map) Set(key, value string) {
	(*c)[key] = value
}

var config *Map

//C gets the global config object
func C() Map {
	return *config
}

//PrettyPrint prints pretty
func (c Map) PrettyPrint() string {
	output := "Key => Value\n"
	for k, v := range c {
		output += fmt.Sprintf("'%s' = '%s'\n", k, v)
	}
	return output
}

//Init initialize config
func Init(envsToRead []string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprint(os.Stderr, errors.Wrap(err, "error in loading env file").Error())
	}
	config = &Map{}
	for _, e := range envsToRead {
		v := os.Getenv(strings.ToUpper(e))
		config.Set(strings.ToLower(e), v)
	}
	Get = config.Get
}
