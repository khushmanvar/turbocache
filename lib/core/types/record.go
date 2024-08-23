package types

type Record struct {
	Value     interface{}
	ExpiresAt int64
}

func NewRecord(val interface{}, exp int64) *Record {
	return &Record{val, exp}
}
