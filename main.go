package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	out, err := toJSON(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, string(out))
}

func toJSON(r io.Reader) ([]byte, error) {
	ybuf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var y interface{}
	if err := yaml.Unmarshal(ybuf, &y); err != nil {
		return nil, err
	}
	y = convert(y)
	j, err := json.Marshal(y)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		strmap := map[string]interface{}{}
		for k, v := range x {
			kstr, ok := k.(string)
			if !ok {
				fmt.Fprintf(os.Stderr, "skipping non-string key %#v with value %#v", k, v)
				continue
			}
			strmap[kstr] = convert(v)
		}
		return strmap
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}
