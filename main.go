package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

var str = `
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
type voidmap map[interface{}]interface{}

func main() {
	m := make(voidmap)

	err := yaml.Unmarshal([]byte(str), &m)
	if err != nil {
		panic(err)
	}

	fmt.Println(m["root"])
	// attempting
	//
	// fmt.Println(m["root"]["foo"])
	//
	// fails b/c indexing not defined for interface{}
	fmt.Println(m["root"].(voidmap)["foo"])
	fmt.Println(m["root"].(voidmap)["baz"])
	fmt.Println(m["root"].(voidmap)["baz"].(voidmap)["omg"])
	fmt.Println(m["root"].(voidmap)["quux"].(voidmap)["hai"].(voidmap)["omg"])
}
