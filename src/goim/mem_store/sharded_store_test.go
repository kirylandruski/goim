package mem_store

import (
	"testing"
	"context"
	"goim/common"
)

func benchmarkSharderdStore(inputs []string, store *SharderStore, limit int, pb *testing.PB) {
	i := 0;
	for pb.Next() {
		store.SetStr(&inputs[i%limit], &inputs[(i+1)%limit], int64(i%10000+1))
		store.GetStr(&inputs[i%limit])
		i++
	}
}

func BenchmarkShardedStoreLockEntryLimited(b *testing.B) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	store := NewShardedStore(context.Background(), 128, false)
	b.RunParallel(func(pb *testing.PB) {
		benchmarkSharderdStore(testInputs, store, testInputsCount, pb)
	})
}

func BenchmarkShardedStoreCowEntryLimited(b *testing.B) {
	const testInputsCount = 10
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	store := NewShardedStore(context.Background(), 128, true)
	b.RunParallel(func(pb *testing.PB) {
		benchmarkSharderdStore(testInputs, store, 10, pb)
	})
}

func BenchmarkShardedStoreLockEntryUnlimited(b *testing.B) {
	const testInputsCount = 10000
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	store := NewShardedStore(context.Background(), 128, false)
	b.RunParallel(func(pb *testing.PB) {
		benchmarkSharderdStore(testInputs, store, testInputsCount, pb)
	})
}

func BenchmarkShardedStoreCowEntryUnlimited(b *testing.B) {
	const testInputsCount = 10000
	const testItemLen = 100
	testInputs := common.GenerateRandomStrings(testInputsCount, testItemLen)

	store := NewShardedStore(context.Background(), 128, true)
	b.RunParallel(func(pb *testing.PB) {
		benchmarkSharderdStore(testInputs, store, testInputsCount, pb)

	})
}
