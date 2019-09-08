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
func Init()error{
	env, err := godotenv.Read(".env")
	if err != nil {
		return errors.Wrap(err, "error in loading env file")
	}
	config = &Map{}
	for k, _ := range env {
		v := os.Getenv(strings.ToUpper(k))
		if v != "" {
			env[k] = v
		}
		config.Set(strings.ToLower(k), env[k])
	}
	Get = config.Get
	return nil
}
