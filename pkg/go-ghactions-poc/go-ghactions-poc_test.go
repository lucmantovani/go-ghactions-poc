package goghactionspoc

import "testing"

func TestHello(t *testing.T) {
	g := GoGHActionsPOC{}
	g.Hello()
	t.Fatal("Hello World!")
}

func TestHello2(t *testing.T) {
	g := GoGHActionsPOC{}
	g.Hello()
	t.Fatal ("Hello 2")
}
