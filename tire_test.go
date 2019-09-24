package tire

import (
	"testing"
)

var t = &Tire{}

func init() {
	t.Insert("/hello/:username/:addr/", "xixi2")
	t.Insert("/hello/:username/:adda/aa", "xixi1")
	t.Insert("/hello/:username/:adda/b", 1)
	t.Insert("/hello/:username/:adda/c", 2)
	t.Insert("/hello/:username/:adda/d", 3)
	t.Insert("/hello/:username/:adda/e", 4)
	t.Insert("/hello/:username/:adda/f", 5)
	t.Insert("/a/:1/2/:2/:2", 6)

}

var p = []byte("/a/1/2/1/ddsadasdsadsadsdsadsaddsadasdsadsadsdsadsa")

func BenchmarkMyTire(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.GetValue(p)
		// t.ParseParams(p)
	}
}

var n1 = new(node)

func init() {
	// n.addRoute("/hello/:username/:addr/", "xixi2")
	n1.addRoute("/hello/:username/:adda/aa", "xixi1")
	n1.addRoute("/hello/:username/:adda/b", 1)
	n1.addRoute("/hello/:username/:adda/c", 2)
	n1.addRoute("/hello/:username/:adda/d", 3)
	n1.addRoute("/hello/:username/:adda/e", 4)
	n1.addRoute("/hello/:username/:adda/f", 5)
	n1.addRoute("/a/:1/2/:2/:2", 6)
}

func BenchmarkTireTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n1.getValue("/a/1/2/1/ddsadasdsadsadsdsadsaddsadasdsadsadsdsadsa")
	}
}