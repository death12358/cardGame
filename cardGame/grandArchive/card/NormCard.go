package card

var Dungeon_Guide = &Card{
	Name: "Dungeon Guide",
}

var Trash_Floating_Memory = &Card{
	Name: "Trash",
	Info: CardInfo{
		Effects: []ICardEffect{Floating_Memory},
	},
}
