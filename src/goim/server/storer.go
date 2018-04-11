package server

type Storer interface {
	GetKeys() ([]*string)
	GetStr(key *string) (res *string, expires int64)
	SetStr(key *string, str *string, ttl int64) (expires int64)
	GetArr(key *string) (res []*string, expires int64)
	SetArr(key *string, arr []*string, ttl int64) (expires int64)
	GetArrItem(key *string, index int32) (res *string, expires int64)
	SetArrItem(key *string, index int32, str *string, ttl int64) (expires int64)
	GetDict(key *string) (res map[string]*string, expires int64)
	SetDict(key *string, dict map[string]*string, ttl int64) (expires int64)
	GetDictItem(key *string, subkey *string) (res *string, expires int64)
	SetDictItem(key *string, subkey *string, str *string, ttl int64) (expires int64)
	RemoveEntry(key *string)
}
