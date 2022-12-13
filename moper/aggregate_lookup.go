package moper

type LU struct {
	d *D
}

// Specifies the foreign collection in the same database to join to the local collection.
// The foreign collection cannot be sharded
func (l *LU) From(collName string) *LU {
	l.d = l.d.Equal("from", collName)
	return l
}

// Specifies the local documents' localField to perform an equality match with the foreign documents' foreignField.
func (l *LU) LocalField(field string) *LU {
	l.d = l.d.Equal("localField", field)
	return l
}

// Specifies the foreign documents' foreignField to perform an equality match with the local documents' localField.
// If a foreign document does not contain a foreignField value, the $lookup uses a null value for the match.
func (l *LU) ForeignField(field string) *LU {
	l.d = l.d.Equal("foreignField", field)
	return l
}

func (l *LU) Custom(p P) *LU {
	l.d = l.d.Equal(p.K, p.V)
	return l
}

func (l *LU) D() *D {
	return NewD().Equal("$lookup", l.d)
}

// Optional, Specifies the variables to use in the pipeline stages.
// Use the variable expressions to access the document fields that are input to the pipeline.
// func (l *LookUp) Let(collName string) *LookUp {
// 	l.d = l.d.Equal("let", collName)
// 	return l
// }

func (l *LU) As(field string) *LU {
	l.d = l.d.Equal("as", field)
	return l
}

func (d *D) LookUp() *LU {
	return &LU{d: d}
}
