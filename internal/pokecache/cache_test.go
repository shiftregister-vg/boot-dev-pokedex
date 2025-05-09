package pokecache_test

import (
	"testing"
	"time"

	"github.com/shiftregister-vg/pokedexcli/internal/pokecache"
)

func TestCache(t *testing.T) {
	key := "test"
	val := []byte("test")

	// create a new cache
	cache := pokecache.NewCache(time.Second * 5)

	// expect no cache hit for key
	bytes, found := cache.Get(key)
	if found || len(bytes) > 0 {
		// test fails
		t.Errorf("expected to find no key with no bytes but found the opposite. Found: %v, Bytes len: %d", found, len(bytes))
		t.Fail()
	}

	// put something in the cache
	cache.Add(key, val)

	// expect found to be true and bytes to match val
	bytes, found = cache.Get(key)
	if !found || len(bytes) == 0 || string(bytes) != string(val) {
		// test fails
		t.Errorf("expected to find a key with bytes but found the opposite. Found: %v, Bytes len: %d", found, len(bytes))
		t.Fail()
	}

	// test the reap
	time.Sleep(time.Second * 15)

	// result should be the same as if the key did not exist
	bytes, found = cache.Get(key)
	if found || len(bytes) > 0 {
		// test fails
		t.Errorf("expected to find no key with no bytes but found the opposite. Found: %v, Bytes len: %d", found, len(bytes))
		t.Fail()
	}

}
