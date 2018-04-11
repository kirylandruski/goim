package mem_store

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// cow stands for copy-on-write
// this implementation doesn't lock read when performing write
type cowEntry struct {
	sync.Mutex
	le unsafe.Pointer
}

func newCowEntry() *cowEntry {
	return &cowEntry{le: unsafe.Pointer(newLockEntry())}
}

func (e *cowEntry) isExpired() bool {
	return e.getLockEntry().isExpired()
}

func (e *cowEntry) getStr() (*string, int64) {
	return e.getLockEntry().getStr()
}

func (e *cowEntry) getArr() ([]*string, int64) {
	return e.getLockEntry().getArr()
}

func (e *cowEntry) getArrItem(index int32) (*string, int64) {
	return e.getLockEntry().getArrItem(index)
}

func (e *cowEntry) getDict() (map[string]*string, int64) {
	return e.getLockEntry().getDict()
}

func (e *cowEntry) getDictItem(subkey *string) (*string, int64) {
	return e.getLockEntry().getDictItem(subkey)
}

func (e *cowEntry) setStr(str *string, ttl int64) int64 {
	e.Lock()
	cp := e.shallowCopy()
	expires := cp.setStr(str, ttl)
	e.setLockEntry(cp)

	e.Unlock()
	return expires
}

func (e *cowEntry) setArr(arr []*string, ttl int64) int64 {
	e.Lock()
	cp := e.shallowCopy()
	expires := cp.setArr(arr, ttl)
	e.setLockEntry(cp)

	e.Unlock()
	return expires
}

func (e *cowEntry) setArrItem(index int32, str *string, ttl int64) int64 {
	e.Lock()
	cp := e.arrCopy()
	expires := cp.setArrItem(index, str, ttl)
	e.setLockEntry(cp)

	e.Unlock()
	return expires
}

func (e *cowEntry) setDict(dict map[string]*string, ttl int64) int64 {
	e.Lock()
	cp := e.shallowCopy()
	expires := cp.setDict(dict, ttl)
	e.setLockEntry(cp)

	e.Unlock()
	return expires
}

func (e *cowEntry) setDictItem(subkey *string, str *string, ttl int64) int64 {
	e.Lock()
	cp := e.dictCopy()
	expires := cp.setDictItem(subkey, str, ttl)
	e.setLockEntry(cp)

	e.Unlock()
	return expires
}

func (e *cowEntry) shallowCopy() *lockEntry {
	le := e.getLockEntry()
	return &lockEntry{
		expires: le.expires,
		str:     le.str,
		arr:     le.arr,
		dict:    le.dict,
	}
}

func (e *cowEntry) arrCopy() *lockEntry {
	le := e.getLockEntry()

	arrCopy := make([]*string, len(le.arr))
	copy(arrCopy, le.arr)
	return &lockEntry{
		expires: le.expires,
		str:     le.str,
		arr:     arrCopy,
		dict:    le.dict,
	}
}

func (e *cowEntry) dictCopy() *lockEntry {
	le := e.getLockEntry()

	dictCopy := make(map[string]*string, len(le.arr))
	for k, v := range le.dict {
		dictCopy[k] = v
	}
	return &lockEntry{
		expires: le.expires,
		str:     le.str,
		arr:     le.arr,
		dict:    dictCopy,
	}
}

func (e *cowEntry) getLockEntry() *lockEntry {
	return (*lockEntry)(atomic.LoadPointer(&e.le))
}

func (e *cowEntry) setLockEntry(entry *lockEntry) {
	atomic.StorePointer(&e.le, unsafe.Pointer(entry))
}
