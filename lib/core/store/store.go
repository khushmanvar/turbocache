package store

import (
	"turbocache/lib/core/types"
	"turbocache/lib/core/utils"
)

var store map[string]*types.Record

func NewRecord(value interface{}, durationMs int64) *types.Record {
	rcd := &types.Record{value, utils.GetExpiresAt(durationMs)}

	return rcd
}

func Get(k string) *types.Record {
	return store[k]
}

func Put(k string, rcd *types.Record) {
	store[k] = rcd
}
