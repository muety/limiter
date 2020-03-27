package memory_test

import (
	"testing"
	"time"

	"github.com/n1try/limiter/v3"
	"github.com/n1try/limiter/v3/drivers/store/memory"
	"github.com/n1try/limiter/v3/drivers/store/tests"
)

func TestMemoryStoreSequentialAccess(t *testing.T) {
	tests.TestStoreSequentialAccess(t, memory.NewStoreWithOptions(limiter.StoreOptions{
		Prefix:          "limiter:memory:sequential",
		CleanUpInterval: 30 * time.Second,
	}))
}

func TestMemoryStoreConcurrentAccess(t *testing.T) {
	tests.TestStoreConcurrentAccess(t, memory.NewStoreWithOptions(limiter.StoreOptions{
		Prefix:          "limiter:memory:concurrent",
		CleanUpInterval: 1 * time.Nanosecond,
	}))
}
