package strategy

import "testing"

func TestStrategy(t *testing.T) {
	cache := initCache(&lfu{})
	cache.add("1", "a")
	cache.add("2", "b")
	cache.add("3", "c")
	if _, ok := cache.get("3"); !ok {
		t.Fatal("key 3 is not exist!")
	}
	cache.setEvictAlgo(&fifo{})
	cache.add("4", "d")
	if _, ok := cache.get("4"); !ok {
		t.Fatal("key 4 is not exist!")
	}
	cache.setEvictAlgo(&lru{})
	cache.add("5", "f")
	if _, ok := cache.get("5"); !ok {
		t.Fatal("key 5 is not exist!")
	}
	cache.del("5")
	if _, ok := cache.get("5"); ok {
		t.Fatal("key 5 is not been deleted")
	}
}
