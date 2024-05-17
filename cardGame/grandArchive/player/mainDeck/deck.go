package mainDeck

import (
	"grandArchive/card"
	"math/rand"
)

type Deck []*card.Card

// 洗牌函数
func shuffle(deck Deck) {
	// 初始化随机数生成器
	// 从最后一个元素开始，逐个与随机位置的元素交换
	for i := len(deck) - 1; i > 0; i-- {
		// 生成一个随机索引 j，其中 0 <= j <= i
		j := rand.Intn(i + 1)

		// 交换 deck[i] 和 deck[j]
		deck[i], deck[j] = deck[j], deck[i]
	}
}
