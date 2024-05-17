package mainDeck

import (
	"fmt"
	"testing"
)

func Test_Shuffle(t *testing.T) {
	// 创建一副卡牌
	deck := Deck{
		{Name: "aaa"},
		{Name: "bbb"},
		{Name: "ccc"},
		{Name: "ddd"},
		// 添加其它卡牌...
	}
	times := 0
	for i := 0; i < 1000000; i++ {

		// 洗牌
		shuffle(deck)
		if deck[0].Name == "aaa" {
			times++
		}

	}
	fmt.Println(times)
}
