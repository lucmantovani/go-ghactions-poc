package goghactionspoc

import (
	"encoding/json"
	"fmt"
)

type GoGHActionsPOC struct{}

func (g *GoGHActionsPOC) Hello() {
	var m string
	json.Unmarshal([]byte("Hello World!"), &m)
	fmt.Println(m)
}
