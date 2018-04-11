package mem_store

import (
	"sync"
	"time"
)

// this implementation of entry users RWMutex and locks all operation when doing writes
type lockEntry struct {
	sync.RWMutex
	expires int64
	str     *string
	arr     []*string
	dict    map[string]*string
}

func newLockEntry() *lockEntry {
	return &lockEntry{expires: 0}
}

func (e *lockEntry) isExpired() bool {
	return e.expires == -1 || (e.expires > 0 && e.expires < (time.Now().UnixNano()/int64(time.Millisecond)))
}

func (e *lockEntry) setTtl(ttl int64) {
	if ttl > 0 {
		e.expires = (time.Now().UnixNano() / int64(time.Millisecond)) + ttl
	} else if ttl == -1 {
		e.expires = 0
	}
}

func (e *lockEntry) clear() {
	e.str = nil
	e.arr = nil
	e.dict = nil
	e.expires = 0
}

func (e *lockEntry) getStr() (*string, int64) {
	e.RLock()
	if e.isExpired() {
		e.RUnlock()
		return nil, -1
	}

	str := e.str
	expires := e.expires
	e.RUnlock()

	return str, expires
}

func (e *lockEntry) getArr() ([]*string, int64) {
	e.RLock()
	if e.isExpired() {
		e.RUnlock()
		return nil, -1
	}

	arr := e.arr
	expires := e.expires
	e.RUnlock()

	return arr, expires
}

func (e *lockEntry) getArrItem(index int32) (*string, int64) {
	e.RLock()
	if e.isExpired() {
		e.RUnlock()
		return nil, -1
	}

	arr := e.arr
	var res *string
	if arr != nil && index < int32(len(arr)) {
		res = arr[index]
	}
	expires := e.expires
	e.RUnlock()

	return res, expires
}

func (e *lockEntry) getDict() (map[string]*string, int64) {
	e.RLock()
	if e.isExpired() {
		e.RUnlock()
		return nil, -1
	}

	dict := e.dict
	expires := e.expires
	e.RUnlock()

	return dict, expires
}

func (e *lockEntry) getDictItem(subkey *string) (*string, int64) {
	e.RLock()
	if e.isExpired() {
		e.RUnlock()
		return nil, -1
	}

	dict := e.dict
	var res *string
	if dict != nil {
		res = dict[*subkey]
	}
	expires := e.expires
	e.RUnlock()

	return res, expires
}

func (e *lockEntry) setStr(str *string, ttl int64) int64 {
	var expires int64

	e.Lock()
	if e.isExpired() {
		e.clear()
	}
	e.setTtl(ttl)
	e.str = str
	expires = e.expires
	e.Unlock()

	return expires
}

func (e *lockEntry) setArr(arr []*string, ttl int64) int64 {
	var expires int64

	e.Lock()
	if e.isExpired() {
		e.clear()
	}
	e.setTtl(ttl)
	e.arr = arr
	expires = e.expires
	e.Unlock()

	return expires
}

func (e *lockEntry) setArrItem(index int32, str *string, ttl int64) int64 {
	var arr []*string
	var expires int64

	e.Lock()
	if e.isExpired() {
		e.clear()
	}
	arr = e.arr
	if arr == nil {
		arr = make([]*string, index+1)
	} else if int32(len(arr)) <= index {
		addition := make([]*string, index-int32(len(arr))+1)
		arr = append(arr, addition...)
	}
	arr[index] = str
	e.setTtl(ttl)
	e.arr = arr
	expires = e.expires
	e.Unlock()

	return expires
}

func (e *lockEntry) setDict(dict map[string]*string, ttl int64) int64 {
	var expires int64

	e.Lock()
	if e.isExpired() {
		e.clear()
	}
	e.setTtl(ttl)
	e.dict = dict
	expires = e.expires
	e.Unlock()

	return expires
}

func (e *lockEntry) setDictItem(subkey *string, str *string, ttl int64) int64 {
	var dict map[string]*string
	var expires int64

	e.Lock()
	if e.isExpired() {
		e.clear()
	}

	dict = e.dict
	if dict == nil {
		dict = make(map[string]*string)
	}
	dict[*subkey] = str
	e.setTtl(ttl)
	e.dict = dict
	expires = e.expires
	e.Unlock()

	return expires
}
