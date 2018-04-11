//go:generate mockgen -source=app_server.go -destination app_server_mock.go --package rpc

package rpc

type AppServer interface {
	Authenticate(username *string, password *string) (status int32)
	GetKeys() (res []*string, status int32)
	GetStr(key *string) (res *string, expires int64, status int32)
	SetStr(key *string, str *string, ttl int64) (expires int64, status int32)
	GetArr(key *string) (res []*string, expires int64, status int32)
	SetArr(key *string, arr []*string, ttl int64) (expires int64, status int32)
	GetArrItem(key *string, index int32) (res *string, expires int64, status int32)
	SetArrItem(key *string, index int32, str *string, ttl int64) (expires int64, status int32)
	GetDict(key *string) (res map[string]*string, expires int64, status int32)
	SetDict(key *string, dict map[string]*string, ttl int64) (expires int64, status int32)
	GetDictItem(key *string, subkey *string) (res *string, expires int64, status int32)
	SetDictItem(key *string, subkey *string, str *string, ttl int64) (expires int64, status int32)
	RemoveEntry(key *string) (status int32)
}
