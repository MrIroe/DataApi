package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

var configObj map[string]string = func() map[string]string {
	b, err := ioutil.ReadFile("config/config.ignore")
	if err != nil {
		return nil
	}
	var locConfObj map[string]string
	err = json.Unmarshal(b, &locConfObj)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Could not read json from config.ignore"))
		return nil
	}
	return locConfObj
}()

func GetConfigValue(val string) string {
	if configObj != nil && configObj[val] != "" {
		return configObj[val]
	} else if os.Getenv(val) != "" {
		return os.Getenv(val)
	} else {
		return ""
	}
}
