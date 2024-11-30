package main

type Node struct {
	key  int
	val  interface{}
	prev *Node
	next *Node
}

type LRUcache struct {
	maxSize int
	size    int
	index   map[int]*Node
	top     *Node
	end     *Node
}

func (c *LRUcache) Set(key int, val interface{}) {
	// if Present
	if val, ok := c.index[key]; ok {
		// update the value
		val.val = val
	} else {
		// else if size 0
		if c.size == 0 {
			// create the first node
			node := Node{
				key:  key,
				val:  val,
				prev: nil,
				next: nil,
			}
			// update size
			c.size += 1
			// update value in map
			c.index[key] = &node
			// update top
			c.top = &node
			// update end
			c.end = &node
			// if size > maxSize
		} else if c.size >= c.maxSize {
			// evict cache
			prev := c.end
			prev.next = nil
			delete(c.index, c.end.key)
			// add new value
			node := Node{
				key:  key,
				val:  val,
				prev: nil,
				next: c.top,
			}
			c.top.prev = &node
			// update map
			c.index[key] = &node
			// update top
			c.top = &node
			// update end
			c.end = c.end.prev
			// else
		} else {
			// set the new val in map
			node := Node{
				key:  key,
				val:  val,
				prev: nil,
				next: c.top,
			}
			c.top.prev = &node
			// update top
			c.top = &node
			// increase current size count
			c.size += 1
			// add the new node to the top
			c.index[key] = &node
		}
	}

}

func (c *LRUcache) Get(key int) interface{} {
	// get Node map
	if val, ok := c.index[key]; ok {
		// if Present
		// if last node -> update the end node
		if val.next == nil {
			prev := val.prev
			prev.next = nil
			val.prev = nil
			c.end = prev
		} else {
			prev := val.prev
			next := val.next
			prev.next = next
			next.prev = prev
		}
		// else bring the node to the top
		val.next = c.top
		c.top.prev = val
		c.top = val
		c.index[key] = val
		// return the val
		return val.val
	} else {
		// else return nil
		return nil
	}
}

func NewLRUCache(maxSize int) *LRUcache {
	cache := LRUcache{
		size:    0,
		maxSize: maxSize,
		index:   make(map[int]*Node),
		top:     nil,
		end:     nil,
	}

	return &cache
}
