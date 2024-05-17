package card

type ICard interface {
}
type IUnit interface {
}

type IAlly interface {
}

type IHero interface {
}

type IDomain interface {
}
type Effect struct{}
type Effect_OnEnter struct{}
type Effect_OnDeath struct{}
type CardEffect struct{}
type ICardEffect interface {
	Activate(target IUnit)
}

func (e Effect) Activate(target IUnit) {

}
