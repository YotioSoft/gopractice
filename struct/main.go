package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	(*p).X = 1e9
	fmt.Println(v)
	p.X = 0				// 本来は(*p).Xだが、p.Xという表記も可
	fmt.Println(v)
}
