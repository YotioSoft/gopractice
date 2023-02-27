package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs" : Vertex{		// Vertexは省略可
		40.68433, -74.39967,
	},
	"Google" : Vertex {
		37.42202, -122.08408,
	},
}
func main() {
	n := make(map[string]Vertex)
	n["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(n["Bell Labs"])

	fmt.Println(m)

	o := make(map[string]int)

	o["Answer"] = 42
	fmt.Println("The value:", o["Answer"])

	o["Answer"] = 48
	fmt.Println("The value:", o["Answer"])

	delete(o, "Answer")
	fmt.Println("The value:", o["Answer"])

	v, ok := o["Answer"]
	fmt.Println("The value:", v, "Present?", ok)	// key"Answer"がないのでokはfalse
}
