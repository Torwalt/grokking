package hashmap

type Node struct {
	key   string
	value int
	next  *Node
}

// Missing: Growing, Delete.

type Map struct {
	buckets []*Node
	size    int
	count   int
}

func NewMap(size int) *Map {
	return &Map{
		buckets: make([]*Node, size),
		size:    size,
	}
}

func (m *Map) Get(key string) (int, bool) {
	idx := m.index(key)

	n := m.buckets[idx]
	if n == nil {
		return 0, false
	}

	return n.value, true
}

func (m *Map) Count() int {
	return m.count
}

func (m *Map) Insert(key string, value int) {
	idx := m.index(key)

	existing := m.buckets[idx]
	n := &Node{
		key:   key,
		value: value,
		next:  nil,
	}

	if existing == nil {
		m.buckets[idx] = n
		m.count++
		return
	}

	if existing.key == key {
		existing.value = value
		return
	}

	existing.next = n
	m.count++
}

func (m *Map) index(key string) int {
	return int(hash(key)) % m.size
}

func hash(key string) uint32 {
	var hash uint32
	for _, ch := range key {
		hash += uint32(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return hash
}
