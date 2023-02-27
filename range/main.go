package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// rangeは、スライスやmapを一つずつ反復処理するのに使う
	// スライスの場合、反復ごとにindexと要素のコピーを返す
	// _で代入を捨てることも可能
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
