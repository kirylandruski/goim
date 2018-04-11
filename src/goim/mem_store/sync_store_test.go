package mem_store

import (
	"testing"
	"reflect"
	"time"
	"context"
	"goim/common"
)

func testSyncStoreStr(store *SyncStore, t *testing.T) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	keys := store.GetKeys()
	if len(keys) != 0 {
		t.Fatal("expected a clean store")
	}

	expires := store.SetStr(&testInputs[0], &testInputs[1], 100)
	if expires < 1 {
		t.Fatal("expected expires to be a date")
	}

	str, expires2 := store.GetStr(&testInputs[0])
	if *str != testInputs[1] {
		t.Fatal("unexpected result")
	}

	expires = store.SetStr(&testInputs[0], &testInputs[5], 0)
	if expires != expires2 {
		t.Fatal("expiration has changed")
	}

	str, expires = store.GetStr(&testInputs[0])
	if *str != testInputs[5] {
		t.Fatal("expected val")
	}

	expires = store.SetStr(&testInputs[0], &testInputs[5], -1)
	if expires != 0 {
		t.Fatal("expiration has changed")
	}

	keys = store.GetKeys()
	if !reflect.DeepEqual(keys, []*string{&testInputs[0]}) {
		t.Fatal("expected to get one key")
	}

	store.RemoveEntry(&testInputs[0])

	str, expires = store.GetStr(&testInputs[0])
	if str != nil || expires != -1 {
		t.Fatal("expected nil")
	}
}

func testSyncStoreArr(store *SyncStore, t *testing.T) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	keys := store.GetKeys()
	if len(keys) != 0 {
		t.Fatal("expected a clean store")
	}

	expires := store.SetArr(&testInputs[1], []*string{&testInputs[0], &testInputs[5], &testInputs[2]}, 0)
	if expires != 0 {
		t.Fatal("unexpected result")
	}

	arr, expires := store.GetArr(&testInputs[1])
	if !reflect.DeepEqual(arr, []*string{&testInputs[0], &testInputs[5], &testInputs[2]}) || expires != 0 {
		t.Fatal("unexpected result")
	}

	expires = store.SetArrItem(&testInputs[2], 2, &testInputs[8], 0)
	if expires != 0 {
		t.Fatal("unexpected result")
	}

	arr, expires = store.GetArr(&testInputs[2])
	if !reflect.DeepEqual(arr, []*string{nil, nil, &testInputs[8]}) {
		t.Fatal("unexpected result")
	}

	str, expires := store.GetArrItem(&testInputs[2], 5)
	if str != nil {
		t.Fatal("unexpected result")
	}

	str, expires = store.GetArrItem(&testInputs[2], 2)
	if *str != testInputs[8] {
		t.Fatal("unexpected result")
	}

	str, expires = store.GetArrItem(&testInputs[3], 2)
	if str != nil || expires != -1 {
		t.Fatal("unexpected result")
	}

	expires = store.SetArr(&testInputs[1], []*string{&testInputs[4], &testInputs[1]}, 0)
	if expires != 0 {
		t.Fatal("unexpected result")
	}

	arr, expires = store.GetArr(&testInputs[1])
	if !reflect.DeepEqual(arr, []*string{&testInputs[4], &testInputs[1]}) || expires != 0 {
		t.Fatal("unexpected result")
	}

	expires = store.SetArrItem(&testInputs[1], 0, &testInputs[3], 0)
	if expires != 0 {
		t.Fatal("unexpected result")
	}

	arr, expires = store.GetArr(&testInputs[1])
	if !reflect.DeepEqual(arr, []*string{&testInputs[3], &testInputs[1]}) || expires != 0 {
		t.Fatal("unexpected result")
	}

	arr, expires = store.GetArr(&testInputs[5])
	if arr != nil || expires != -1 {
		t.Fatal("unexpected result")
	}

	keys = store.GetKeys()
	if len(keys) != 2 {
		t.Fatal("two keys")
	}

	if *keys[0] != testInputs[1] && *keys[1] != testInputs[1] {
		t.Fatal("two keys")
	}

	if *keys[0] != testInputs[2] && *keys[1] != testInputs[2] {
		t.Fatal("two keys")
	}
}

func testSyncStoreDict(store *SyncStore, t *testing.T) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	keys := store.GetKeys()
	if len(keys) != 0 {
		t.Fatal("expected a clean store")
	}

	expires := store.SetDict(&testInputs[1], map[string]*string{testInputs[0]: &testInputs[0], testInputs[1]: &testInputs[5], testInputs[2]: &testInputs[2]}, 0)
	if expires != 0 {
		t.Fatal("unexpected result")
	}

	dict, expires := store.GetDict(&testInputs[1])
	if !reflect.DeepEqual(dict, map[string]*string{testInputs[0]: &testInputs[0], testInputs[1]: &testInputs[5], testInputs[2]: &testInputs[2]}) || expires != 0 {
		t.Fatal("unexpected result")
	}

	expires = store.SetDictItem(&testInputs[2], &testInputs[4], &testInputs[8], 0)
	if expires != 0 {
		t.Fatal("unexpected result")
	}

	dict, expires = store.GetDict(&testInputs[2])
	if !reflect.DeepEqual(dict, map[string]*string{testInputs[4]: &testInputs[8]}) {
		t.Fatal("unexpected result")
	}

	str, expires := store.GetDictItem(&testInputs[2], &testInputs[9])
	if str != nil {
		t.Fatal("unexpected result")
	}

	str, expires = store.GetDictItem(&testInputs[2], &testInputs[4])
	if *str != testInputs[8] {
		t.Fatal("unexpected result")
	}

	str, expires = store.GetDictItem(&testInputs[3], &testInputs[4])
	if str != nil || expires != -1 {
		t.Fatal("unexpected result")
	}

	expires = store.SetDict(&testInputs[1], map[string]*string{testInputs[3]: &testInputs[9]}, 0)
	if expires != 0 {
		t.Fatal("unexpected result")
	}

	dict, expires = store.GetDict(&testInputs[1])
	if !reflect.DeepEqual(dict, map[string]*string{testInputs[3]: &testInputs[9]}) || expires != 0 {
		t.Fatal("unexpected result")
	}

	expires = store.SetDictItem(&testInputs[1], &testInputs[1], &testInputs[4], 0)
	if expires != 0 {
		t.Fatal("unexpected result")
	}

	dict, expires = store.GetDict(&testInputs[1])
	if !reflect.DeepEqual(dict, map[string]*string{testInputs[3]: &testInputs[9], testInputs[1]: &testInputs[4]}) || expires != 0 {
		t.Fatal("unexpected result")
	}

	dict, expires = store.GetDict(&testInputs[5])
	if dict != nil || expires != -1 {
		t.Fatal("unexpected result")
	}

	keys = store.GetKeys()
	if len(keys) != 2 {
		t.Fatal("two keys")
	}

	if *keys[0] != testInputs[1] && *keys[1] != testInputs[1] {
		t.Fatal("two keys")
	}

	if *keys[0] != testInputs[2] && *keys[1] != testInputs[2] {
		t.Fatal("two keys")
	}
}

func testSyncStoreStrExpiration(store *SyncStore, t *testing.T) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	keys := store.GetKeys()
	if len(keys) != 0 {
		t.Fatal("expected a clean store")
	}

	store.SetArr(&testInputs[0], []*string{&testInputs[0], &testInputs[5], &testInputs[2]}, 1)
	time.Sleep(time.Second * 2)
	expires := store.SetStr(&testInputs[0], &testInputs[8], 0)
	if expires != 0 {
		t.Fatal("as long as old item expired this should be a zero")
	}

	arr, expires := store.GetArr(&testInputs[0])
	if arr != nil {
		t.Fatal("expected nil str")
	}
	str, expires := store.GetStr(&testInputs[0])
	if *str != testInputs[8] {
		t.Fatal("expected str")
	}
}

func testSyncStoreArrExpiration(store *SyncStore, t *testing.T) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	keys := store.GetKeys()
	if len(keys) != 0 {
		t.Fatal("expected a clean store")
	}

	store.SetStr(&testInputs[0], &testInputs[1], 1)
	time.Sleep(time.Second * 2)
	expires := store.SetArr(&testInputs[0], []*string{&testInputs[0], &testInputs[5], &testInputs[2]}, 0)
	if expires != 0 {
		t.Fatal("as long as old item expired this should be a zero")
	}

	str, expires := store.GetStr(&testInputs[0])
	if str != nil {
		t.Fatal("expected nil str")
	}
	arr, expires := store.GetArr(&testInputs[0])
	if !reflect.DeepEqual(arr, []*string{&testInputs[0], &testInputs[5], &testInputs[2]}) {
		t.Fatal("expected str")
	}
}

func testSyncStoreArrItemExpiration(store *SyncStore, t *testing.T) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	keys := store.GetKeys()
	if len(keys) != 0 {
		t.Fatal("expected a clean store")
	}

	store.SetStr(&testInputs[0], &testInputs[1], 1)
	time.Sleep(time.Second * 2)
	expires := store.SetArrItem(&testInputs[0], 2, &testInputs[2], 0)
	if expires != 0 {
		t.Fatal("as long as old item expired this should be a zero")
	}

	str, expires := store.GetStr(&testInputs[0])
	if str != nil {
		t.Fatal("expected nil str")
	}
	str, expires = store.GetArrItem(&testInputs[0], 2)
	if *str != testInputs[2] {
		t.Fatal("expected str")
	}
}

func testSyncStoreDictExpiration(store *SyncStore, t *testing.T) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	keys := store.GetKeys()
	if len(keys) != 0 {
		t.Fatal("expected a clean store")
	}

	store.SetStr(&testInputs[0], &testInputs[1], 1)
	time.Sleep(time.Second * 2)
	expires := store.SetDict(&testInputs[0], map[string]*string{testInputs[0]: &testInputs[0], testInputs[5]: &testInputs[2]}, 0)
	if expires != 0 {
		t.Fatal("as long as old item expired this should be a zero")
	}

	str, expires := store.GetStr(&testInputs[0])
	if str != nil {
		t.Fatal("expected nil str")
	}
	arr, expires := store.GetDict(&testInputs[0])
	if !reflect.DeepEqual(arr, map[string]*string{testInputs[0]: &testInputs[0], testInputs[5]: &testInputs[2]}) {
		t.Fatal("expected str")
	}
}

func testSyncStoreDictItemExpiration(store *SyncStore, t *testing.T) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	keys := store.GetKeys()
	if len(keys) != 0 {
		t.Fatal("expected a clean store")
	}

	store.SetStr(&testInputs[0], &testInputs[1], 1)
	time.Sleep(time.Second * 2)
	expires := store.SetDictItem(&testInputs[0], &testInputs[1], &testInputs[2], 0)
	if expires != 0 {
		t.Fatal("as long as old item expired this should be a zero")
	}

	str, expires := store.GetStr(&testInputs[0])
	if str != nil {
		t.Fatal("expected nil str")
	}
	str, expires = store.GetDictItem(&testInputs[0], &testInputs[1])
	if *str != testInputs[2] {
		t.Fatal("expected str")
	}
}

func testSyncStoreGetExpiredItem(store *SyncStore, t *testing.T) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	store.SetStr(&testInputs[0], &testInputs[1], 1)
	store.SetDict(&testInputs[0], map[string]*string{testInputs[0]: &testInputs[0], testInputs[5]: &testInputs[2]}, 0)
	store.SetArr(&testInputs[0], []*string{&testInputs[0], &testInputs[5], &testInputs[2]}, 0)

	time.Sleep(time.Second * 2)

	str, _ := store.GetStr(&testInputs[0])
	if str != nil {
		t.Fatal("expected nil")
	}
	arr, _ := store.GetArr(&testInputs[0])
	if arr != nil {
		t.Fatal("expected nil")
	}
	str, _ = store.GetArrItem(&testInputs[0], 0)
	if str != nil {
		t.Fatal("expected nil")
	}
	dict, _ := store.GetDict(&testInputs[0])
	if dict != nil {
		t.Fatal("expected nil")
	}
	str, _ = store.GetDictItem(&testInputs[0], &testInputs[0])
	if str != nil {
		t.Fatal("expected nil")
	}
}

func testSyncStore(cow bool, t *testing.T) {
	store := NewSyncStore(context.Background(), cow)
	testSyncStoreStr(store, t)

	store = NewSyncStore(context.Background(), cow)
	testSyncStoreArr(store, t)

	store = NewSyncStore(context.Background(), cow)
	testSyncStoreDict(store, t)

	store = NewSyncStore(context.Background(), cow)
	testSyncStoreStrExpiration(store, t)

	store = NewSyncStore(context.Background(), cow)
	testSyncStoreArrExpiration(store, t)

	store = NewSyncStore(context.Background(), cow)
	testSyncStoreArrItemExpiration(store, t)

	store = NewSyncStore(context.Background(), cow)
	testSyncStoreDictExpiration(store, t)

	store = NewSyncStore(context.Background(), cow)
	testSyncStoreDictItemExpiration(store, t)

	store = NewSyncStore(context.Background(), cow)
	testSyncStoreGetExpiredItem(store, t)
}

func TestSyncStoreLockEntry(t *testing.T) {
	testSyncStore(false, t)
}

func TestSyncStoreCowEntry(t *testing.T) {
	testSyncStore(true, t)
}

func benchmarkSyncStore(inputs []string, store *SyncStore, limit int, pb *testing.PB) {
	i := 0;
	for pb.Next() {
		store.SetStr(&inputs[i%limit], &inputs[(i+1)%limit], int64(i%10000+1))
		store.GetStr(&inputs[i%limit])
		i++
	}
}

func BenchmarkSyncStoreLockEntryLimited(b *testing.B) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	store := NewSyncStore(context.Background(), false)
	b.RunParallel(func(pb *testing.PB) {
		benchmarkSyncStore(testInputs, store, testInputsCount, pb)
	})
}

func BenchmarkSyncStoreCowEntryLimited(b *testing.B) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	store := NewSyncStore(context.Background(), true)
	b.RunParallel(func(pb *testing.PB) {
		benchmarkSyncStore(testInputs, store, testInputsCount, pb)
	})
}

func BenchmarkSyncStoreLockEntryUnlimited(b *testing.B) {
	const testInputsCount = 100000
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	store := NewSyncStore(context.Background(), false)
	b.RunParallel(func(pb *testing.PB) {
		benchmarkSyncStore(testInputs, store, testInputsCount, pb)
	})
}

func BenchmarkSyncStoreCowEntryUnlimited(b *testing.B) {
	const testInputsCount = 100000
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	store := NewSyncStore(context.Background(), true)
	b.RunParallel(func(pb *testing.PB) {
		benchmarkSyncStore(testInputs, store, testInputsCount, pb)
	})
}
