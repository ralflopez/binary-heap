package queue

type MinHeap []*Node

func (m MinHeap) ParentIndex(i int) int {
	if i == 0 {
		return -1
	}

	return (i - 1) / 2
}

func (m MinHeap) MinChildIndex(i int) int {
	if len(m) == 0 {
		return -1
	}

	offset := i
	l := (i + 1) + offset
	r := (i + 2) + offset

	if !m.IsValidIndex(l) {
		return -1
	}

	if m.IsValidIndex(l) && !m.IsValidIndex(r) {
		return l
	}

	left := m[l]
	right := m[r]
	if left.Value < right.Value {
		return l
	}

	return r
}

func (m MinHeap) Root() *Node {
	if len(m) == 0 {
		return nil
	}

	return m[0]
}

func (m MinHeap) IsFull() bool {
	return len(m) == cap(m)
}

func (m MinHeap) IsEmpty() bool {
	return len(m) == 0
}

func (m MinHeap) IsValidIndex(i int) bool {
	if i < 0 {
		return false
	}

	if i > len(m)-1 {
		return false
	}

	return true
}

func (m MinHeap) swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Insert(value *Node) {
	if m.IsEmpty() {
		*m = append(*m, value)
		return
	}

	if m.IsFull() && value.Value <= m.Root().Value {
		return
	}

	if m.IsFull() {
		m.Pop()
	}

	*m = append(*m, value)

	c := len(*m) - 1
	for p := m.ParentIndex(c); p != -1; {
		parent := (*m)[p]
		cur := (*m)[c]
		if cur.Value < parent.Value {
			m.swap(p, c)
		}
		c--
		p = m.ParentIndex(c)
	}
}

func (m *MinHeap) Pop() *Node {
	if m.IsEmpty() {
		return nil
	}

	root := (*m)[0]

	n := len(*m) - 1
	(*m)[0] = (*m)[n]
	*m = (*m)[:n]

	c := 0
	minIdx := m.MinChildIndex(c)
	for minIdx != -1 {
		max := (*m)[minIdx]
		cur := (*m)[c]

		if cur.Value <= max.Value {
			break
		}

		m.swap(minIdx, c)
		c++
		minIdx = m.MinChildIndex(c)
	}

	return root
}

func (m MinHeap) ToArray() []*Node {
	arr := make([]*Node, 0, len(m))

	for len(m) > 0 {
		item := m.Pop()
		arr = append(arr, item)
	}

	return arr
}
