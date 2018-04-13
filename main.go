package main

import (
	"math/rand"
	"time"
)

func main() {
	// 現在時刻を数値で取得する
	t := time.Now().UnixNano()
	// 乱数のたねを設定
	rand.Seed(t)
	// xは0-10の間の値になる
	s := rand.Intn(5)

	switch s {
		case 1:
			println("凶")
		case 2, 3:
			println("吉")
		case 4, 5:
			println("中吉")
		case 0:
			println("大吉")

	}
}
