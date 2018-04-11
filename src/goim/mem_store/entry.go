package mem_store

type entry interface {
	Lock()
	Unlock()

	isExpired() bool
	getStr() (*string, int64)
	getArr() ([]*string, int64)
	getArrItem(index int32) (*string, int64)
	getDict() (map[string]*string, int64)
	getDictItem(subkey *string) (*string, int64)
	setStr(str *string, ttl int64) int64
	setArr(arr []*string, ttl int64) int64
	setArrItem(index int32, str *string, ttl int64) int64
	setDict(dict map[string]*string, ttl int64) int64
	setDictItem(subkey *string, str *string, ttl int64) int64
}
