package card

var Creative_Shock = &Card{
	Name:  "Creative_Shock",
	Types: []CardType{Fire},
	Info: CardInfo{
		Cost:    3,
		Effects: []ICardEffect{},
	},
}

func a() {
	Creative_Shock.Info.Effects[0].Activate = 
}
