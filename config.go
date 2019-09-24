package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

var config = &Map{}

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

func envParser(lines []string) map[string]string {
	env := make(map[string]string)
	for _, line := range lines {
		tempLineSplitted := strings.Split(line, "=")
		env[tempLineSplitted[0]] = tempLineSplitted[1]
	}
	return env
}

//Init initialize config
func Init() {
	envBytes, err := ioutil.ReadFile(".env")
	if err != nil {
		fmt.Fprintf(os.Stdout, "erorr in reading env file: %v", err)
		os.Exit(1)
	}
	envLines := strings.Split(string(envBytes), "\n")
	env := envParser(envLines)
	config = &Map{}
	for k, _ := range env {
		v := os.Getenv(strings.ToUpper(k))
		if v != "" {
			env[k] = v
		}
		config.Set(strings.ToLower(k), env[k])
	}
	fmt.Println(config)
	Get = config.Get
}

func main() {
	Init()
}
