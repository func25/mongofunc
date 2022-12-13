package moper

func (d *D) Match(pairs ...P) *D {
	return d.Equal("$match", toPair(pairs))
}

func (d *D) MatchD(pair D) *D {
	return d.Equal("$match", pair)
}

func (d *D) Group(pairs ...P) *D {
	return d.Equal("$group", toPair(pairs))
}

func (d *D) GroupD(pair D) *D {
	return d.Equal("$group", pair)
}

func (d *D) Sum(fieldName string) *D {
	return d.Equal("$sum", "$"+fieldName)
}
