package player

import (
	"grandArchive/card"
)

type Player struct {
	Hero        *card.Card
	Hands       []*card.Card
	DamageTaken int
}

type ICardSet interface {
	HasCard(name string)
	HasCard_withEffect(e card.ICardEffect)
}
type IPlayer interface {
	searchCard(n int)
	drawCard(n int)
	discard(n int, cards []card.Card)
	discard_random(n int)
}

func (p Player) isLose() bool {
	if p.Hero.Info.Hp < p.DamageTaken {
		return true
	}
	return false
}
