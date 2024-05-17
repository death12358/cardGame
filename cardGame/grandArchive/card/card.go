package card

type Card struct {
	Types []CardType
	Name  string
	Info  CardInfo
}
type SpinRequest struct {
	Bet                    int32
	GameCode               string
	PlayerID               int
	CountryID              int
	VendorID               int
	PlatformID             int
	GameRequestTriggerType int
}

type CardType string
type CardInfo struct {
	Cost    int
	Atk     int
	Hp      int
	Effects []ICardEffect
}

func (e Effect_OnEnter) Activate(target IUnit) {}
