package mem_store

import (
	"hash/fnv"
	"golang.org/x/net/context"
)

type SharderStore struct {
	shards []*SyncStore
}

// Implementation of Storer which uses a set of SyncStore and shares load across them basing on a key hash
func NewShardedStore(ctx context.Context, shardsCount uint32, cow bool) *SharderStore {
	s := SharderStore{}
	s.shards = make([]*SyncStore, shardsCount)
	for i := uint32(0); i < shardsCount; i++ {
		s.shards[i] = NewSyncStore(ctx, cow)
	}

	return &s
}

func (s *SharderStore) GetKeys() []*string {
	keys := make([]*string, 0)
	for i := 0; i < len(s.shards); i++ {
		keys = append(keys, s.shards[i].GetKeys()...)
	}

	return keys
}

func (s *SharderStore) getShard(key *string) *SyncStore {
	hash := fnv.New64a()
	hash.Write([]byte(*key))
	n := hash.Sum64() % uint64(len(s.shards))

	return s.shards[n];
}

func (s *SharderStore) GetStr(key *string) (*string, int64) {
	sh := s.getShard(key)
	return sh.GetStr(key)
}

func (s *SharderStore) SetStr(key *string, str *string, ttl int64) int64 {
	sh := s.getShard(key)
	return sh.SetStr(key, str, ttl)
}

func (s *SharderStore) GetArr(key *string) ([]*string, int64) {
	sh := s.getShard(key)
	return sh.GetArr(key)
}

func (s *SharderStore) SetArr(key *string, arr []*string, ttl int64) int64 {
	sh := s.getShard(key)
	return sh.SetArr(key, arr, ttl)
}

func (s *SharderStore) GetArrItem(key *string, index int32) (*string, int64) {
	sh := s.getShard(key)
	return sh.GetArrItem(key, index)
}

func (s *SharderStore) SetArrItem(key *string, index int32, str *string, ttl int64) int64 {
	sh := s.getShard(key)
	return sh.SetArrItem(key, index, str, ttl)
}

func (s *SharderStore) GetDict(key *string) (map[string]*string, int64) {
	sh := s.getShard(key)
	return sh.GetDict(key)
}

func (s *SharderStore) SetDict(key *string, dict map[string]*string, ttl int64) int64 {
	sh := s.getShard(key)
	return sh.SetDict(key, dict, ttl)
}

func (s *SharderStore) GetDictItem(key *string, subkey *string) (*string, int64) {
	sh := s.getShard(key)
	return sh.GetDictItem(key, subkey)
}

func (s *SharderStore) SetDictItem(key *string, subkey *string, str *string, ttl int64) int64 {
	sh := s.getShard(key)
	return sh.SetDictItem(key, subkey, str, ttl)
}

func (s *SharderStore) RemoveEntry(key *string) {
	sh := s.getShard(key)
	sh.RemoveEntry(key)
}
