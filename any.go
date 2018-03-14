package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

var str0 = `
root:
  foo: 0
  bar: 1
  baz:
    omg: 3
    wtf: 4
    lol: 5
  quux:
    hai:
      omg: 6
      wtf: 7
      lol: 8
`

var str1 = `
[0 2 3]
`

func main() {
	var doc interface{}
	err := yaml.Unmarshal([]byte(str1), &doc)
	if err != nil {
		panic(err)
	}

	m := doc.(map[interface{}]interface{})
	fmt.Println(m["root"])
	fmt.Println(m["root"].(map[interface{}]interface{})["foo"])
}
