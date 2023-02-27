package main

import "fmt"

func main() {
	// スライスの基本
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	// スライスは配列の参照のようなもの
	// スライスの要素を変更すると、元の配列の要素も変更される
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// スライスのリテラルは長さのない配列リテラルのようなもの
	// 配列リテラル：[3]bool{true, true, false}
	// スライス：[]bool{true, true, false}
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(t)

	// スライスの長さ(length)と容量(capacity)
	printSlice(q)

	q = q[:0]
	printSlice(q)

	q = q[:4]
	printSlice(q)

	q = q[2:]
	printSlice(q)

	// make関数でスライスを作成
	c := make([]int, 5)
	printSlice(c)

	d := make([]int, 0, 5)
	printSlice(d)

	e := d[:2]
	printSlice(e)

	f := e[2:5]
	printSlice(f)

	// スライスへ新しい要素を追加
	var u []int
	printSlice(u)

	// append works on nil slices.
	u = append(u, 0)
	printSlice(u)

	// The slice grows as needed.
	u = append(u, 1)
	printSlice(u)

	// We can add more then one element at a time.
	u = append(u, 2, 3, 4)
	printSlice(u)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
