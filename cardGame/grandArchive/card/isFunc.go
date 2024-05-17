package card

func (c Card) isCaed(name string) bool {
	if c.Name == name {
		return true
	}
	return false
}

func (c Card) isType(t1 CardType) bool {
	for _, t := range c.Types {
		if t == t1 {
			return true
		}
	}
	return false
}
