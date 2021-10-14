package strategy

import "fmt"

// 策略接口
type evictionAlgo interface {
	evict(c *cache)
}

// fifo策略
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strategy")
}

// lru 策略
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}

// lfu 策略
type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strategy")
}

// cache 上下文角色
type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// 客户端代码

func RunApplication() {

	lfux := &lfu{}
	c := initCache(lfux)

	c.add("a", "b")
	c.add("b", "b")
	c.add("c", "b")

	fifox := &fifo{}
	c.setEvictionAlgo(fifox)
	c.add("d", "b")

	lrux := &lru{}
	c.setEvictionAlgo(lrux)
	c.add("e", "b")
}
