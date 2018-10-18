package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	o := flag.String("o", "", "Write output to this file instead of stdout")
	flag.Parse()

	// If there are more than 1 positional arguments, raise error.
	if flag.NArg() > 1 {
		flag.Usage()
		os.Exit(1)
	}

	// If there are no position arguments, then read from STDIN.
	// If there is one positional argument, then read from that file.
	input := os.Stdin
	if flag.NArg() == 1 {
		var err error
		input, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	out, err := toJSON(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	w := os.Stdout
	if *o != "" {
		w, err = os.OpenFile(*o, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	fmt.Fprint(w, string(out))
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
