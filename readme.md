nested yaml in golang
=====================

Messing around in golang, I was having some difficulty getting a clear
explanation of how to parse an arbitrarily-shaped json/yaml doc, and
then pull out keys, given that the natural way to parse arbitrary
input is as `map[interface{}]interface{}` (or possibly
`map[string]interface{}`, but whatever).

Yeah, I know this isn't how you're supposed to write Go. Leave me
alone. I'M LEARNING THINGS.

- `main.go`: first version I came up with

- `any.go`: much uglier version required to both permit reading
  arbitrary docs (yaml/json doesn't have to be a map) _and_ comply
  with golang's type identity rules.
