package goghactionspoc

import (
	"encoding/json"
	"fmt"
)

type GoGHActionsPOC struct{}

func (g *GoGHActionsPOC) Hello() {
	var m string
	err := json.Unmarshal([]byte("Hello World!"), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}

func (g *GoGHActionsPOC) World() {
	var m string
	err := json.Unmarshal([]byte("Hello World!"), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}
