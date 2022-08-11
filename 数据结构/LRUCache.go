package main

import "container/list"

type LRUCache struct {
	capacity int
	// 双向链表
	ll *list.List
	// 缓存
	cache map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.cache[key]; ok {
		// 移动到头部
		this.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.cache[key]; ok { // 如果存在，就覆盖掉
		this.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		kv.value = value
	} else { // 不存在
		ele := this.ll.PushFront(&entry{key, value})
		this.cache[key] = ele
	}

	if this.capacity < len(this.cache) {
		this.removeOldest()
	}
}

func (this *LRUCache) removeOldest() {
	ele := this.ll.Back()
	if ele != nil {
		this.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(this.cache, kv.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
