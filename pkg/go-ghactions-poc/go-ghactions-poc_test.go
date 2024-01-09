package goghactionspoc

import "testing"

func TestHello(t *testing.T) {
	g := GoGHActionsPOC{}
	g.Hello()
	t.Fatal("Hello World!")
}
