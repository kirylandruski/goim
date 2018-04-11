package mem_store

import (
	"sync"
	"time"
	"math"
	"golang.org/x/net/context"
)

type SyncStore struct {
	root    sync.Map
	factory func() entry
}

func newEntryFactory(cow bool) func() entry {
	if cow {
		return func() entry { return newCowEntry() }
	} else {
		return func() entry { return newLockEntry() }
	}
}

func NewSyncStore(ctx context.Context, cow bool) *SyncStore {
	entryFactory := newEntryFactory(cow)
	store := &SyncStore{factory: entryFactory, root: sync.Map{}}

	go store.startCleanupWorker(ctx)

	return store
}

func (s *SyncStore) GetKeys() []*string {
	keys := make([]*string, 0, 0)
	s.root.Range(func(key, value interface{}) bool {
		str := (key).(*string)
		keys = append(keys, str)
		return true
	})

	return keys
}

func (s *SyncStore) GetStr(key *string) (*string, int64) {
	e := s.getEntry(key)
	if e != nil {
		return e.getStr()
	}

	return nil, -1
}

func (s *SyncStore) GetArr(key *string) ([]*string, int64) {
	e := s.getEntry(key)
	if e != nil {
		return e.getArr()
	}

	return nil, -1
}

func (s *SyncStore) GetArrItem(key *string, index int32) (*string, int64) {
	e := s.getEntry(key)
	if e != nil {
		return e.getArrItem(index)
	}

	return nil, -1
}

func (s *SyncStore) GetDict(key *string) (map[string]*string, int64) {
	e := s.getEntry(key)
	if e != nil {
		return e.getDict()
	}

	return nil, -1
}

func (s *SyncStore) GetDictItem(key *string, subkey *string) (*string, int64) {
	e := s.getEntry(key)
	if e != nil {
		return e.getDictItem(subkey)
	}

	return nil, -1
}

func (s *SyncStore) SetStr(key *string, str *string, ttl int64) int64 {
	e := s.getEntry(key)
	if e != nil {
		expires := e.setStr(str, ttl)
		return expires
	} else {
		e := s.factory()
		expires := e.setStr(str, ttl)
		s.setEntry(key, e)
		return expires
	}
}

func (s *SyncStore) SetArr(key *string, arr []*string, ttl int64) int64 {
	e := s.getEntry(key)
	if e != nil {
		return e.setArr(arr, ttl)
	} else {
		e := s.factory()
		expires := e.setArr(arr, ttl)
		s.setEntry(key, e)

		return expires
	}
}

func (s *SyncStore) SetArrItem(key *string, index int32, str *string, ttl int64) int64 {
	e := s.getEntry(key)
	if e != nil {
		return e.setArrItem(index, str, ttl)
	} else {
		e := s.factory()
		expires := e.setArrItem(index, str, ttl)
		s.setEntry(key, e)

		return expires
	}
}

func (s *SyncStore) SetDict(key *string, dict map[string]*string, ttl int64) int64 {
	e := s.getEntry(key)
	if e != nil {
		return e.setDict(dict, ttl)
	} else {
		e := s.factory()
		expires := e.setDict(dict, ttl)
		s.setEntry(key, e)

		return expires
	}
}

func (s *SyncStore) SetDictItem(key *string, subkey *string, str *string, ttl int64) int64 {
	e := s.getEntry(key)
	if e != nil {
		return e.setDictItem(subkey, str, ttl)
	} else {
		e := s.factory()
		expires := e.setDictItem(subkey, str, ttl)
		s.setEntry(key, e)

		return expires
	}
}

func (s *SyncStore) RemoveEntry(key *string) {
	s.removeEntry(key)
}

func (s *SyncStore) getEntry(key *string) entry {
	e, ok := s.root.Load(key)
	if ok {
		return e.(entry)
	}
	return nil
}

func (s *SyncStore) setEntry(key *string, e entry) {
	s.root.Store(key, e)
}

func (s *SyncStore) removeEntry(key *string) {
	s.root.Delete(key)
}

func (s *SyncStore) startCleanupWorker(ctx context.Context) {
	// the perfect rate of expired/total items
	perfectRate := float64(0.25)

	maxDuration := time.Second
	// initial duration of delay before next cleanup
	duration := 50 * time.Millisecond

	for {
		select {
		case <-ctx.Done():
			return
		default:
			removed, total := s.removeExpired()

			if removed == 0 || total == 0 {
				// set max delay time
				duration = maxDuration
			} else {
				// actual rate of expired/total items
				rate := float64(removed) / float64(total)

				// if actual rate is more than perfect rate ( like 0.25 / 0.35 ) then factor will be less than zero
				// if actual rate is less than perfect rate (like 0.25 / 0.1 ) then factor will be more than zero
				factor := math.Pow(perfectRate/rate, 2)
				// based on the factor delay duration is increased or decreased
				// so if the load is not changed
				duration = time.Duration(int64(factor * float64(duration.Nanoseconds())))
				if duration > maxDuration {
					duration = maxDuration
				}
			}

			time.Sleep(duration)
		}
	}

}

func (s *SyncStore) removeExpired() (removed, total int64) {
	s.root.Range(func(key, value interface{}) bool {
		total ++
		v := (value).(entry)
		v.Lock()
		if v.isExpired() {
			removed++
			s.root.Delete(key)
		}
		v.Unlock()
		return true
	})

	return
}
