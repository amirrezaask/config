package athena

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

//Cfg is a placeholder for config
type Cfg struct {
	m map[string]string
}

func (c *Cfg) Get(key string) string {
	val, exists := (c.m)[key]
	if !exists {
		return ""
	}
	return val
}
func (c *Cfg) Set(key string, val string) {
	c.m[key] = val
}

//PrettyPrint prints pretty
func (c *Cfg) PrettyPrint() string {
	output := "Key => Value\n"
	for k, v := range c.m {
		output += fmt.Sprintf("'%s' = '%s'\n", k, v)
	}
	return output
}

//Init initialize config
func Init(r io.Reader) *Cfg {
	dotEnv, err := godotenv.Parse(r)
	if err != nil {
		fmt.Fprint(os.Stderr, errors.Wrap(err, "error in loading env file").Error())
	}
	for k, _ := range dotEnv {
		env := os.Getenv(strings.ToUpper(k))
		if env != "" {
			dotEnv[k] = env
		}
	}
	return &Cfg{
		dotEnv,
	}
}
