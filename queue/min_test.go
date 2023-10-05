package queue

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("should insert if queue is empty", func(t *testing.T) {
		m := make(MinHeap, 0, 10)
		m.Insert(mockNode(1))

		if m[0].Value != 1 {
			t.Fatalf("failed to insert if queue is empty")
		}
	})

	t.Run("should swap root with left child", func(t *testing.T) {
		m := make(MinHeap, 0, 10)

		input := []float64{2, 1}
		for _, v := range input {
			m.Insert(mockNode(v))
		}

		expected := []float64{1, 2}
		assertEquals(t, m, expected)
	})

	t.Run("should swap root with right child", func(t *testing.T) {
		m := make(MinHeap, 0, 10)

		input := []float64{3, 4, 2}
		for _, v := range input {
			m.Insert(mockNode(v))
		}

		expected := []float64{2, 4, 3}
		assertEquals(t, m, expected)
	})

	t.Run("should bubble up from left child until root", func(t *testing.T) {
		m := make(MinHeap, 0, 10)

		input := []float64{3, 4, 2, 1}
		for _, v := range input {
			m.Insert(mockNode(v))
		}

		expected := []float64{1, 2, 3, 4}
		assertEquals(t, m, expected)
	})

	t.Run("should bubble up from right child until root", func(t *testing.T) {
		m := make(MinHeap, 0, 10)

		input := []float64{5, 6, 7, 8, 4}
		for _, v := range input {
			m.Insert(mockNode(v))
		}

		expected := []float64{4, 5, 7, 8, 6}
		assertEquals(t, m, expected)
	})

	t.Run("should bubble up from left child until invariant", func(t *testing.T) {
		m := make(MinHeap, 0, 10)

		input := []float64{1, 4, 5, 7, 8, 9, 10, 2}
		for _, v := range input {
			m.Insert(mockNode(v))
		}

		expected := []float64{1, 2, 5, 4, 8, 9, 10, 7}
		assertEquals(t, m, expected)
	})

	t.Run("should bubble up from right child until invariant", func(t *testing.T) {
		m := make(MinHeap, 0, 10)

		input := []float64{1, 4, 5, 7, 8, 9, 10, 11, 2}
		for _, v := range input {
			m.Insert(mockNode(v))
		}

		expected := []float64{1, 2, 5, 4, 8, 9, 10, 11, 7}
		assertEquals(t, m, expected)
	})

	t.Run("should not insert if max capacity reached and new value is less than root", func(t *testing.T) {
		maxCap := 1
		m := make(MinHeap, 0, maxCap)
		m.Insert(mockNode(10))

		m.Insert(mockNode(1))

		if len(m) != 1 {
			t.Fatalf("expected %v but got %v", 1, len(m))
		}
	})

	t.Run("  capacity reached and new value is greater than root", func(t *testing.T) {
		maxCap := 5
		m := mockMinHeap([]float64{1, 2, 3, 4, 5})

		m.Insert(mockNode(6))

		if len(m) > maxCap {
			t.Fatalf("expected %v but got %v", maxCap, len(m))
		}

		expected := []float64{2, 4, 3, 5, 6}
		assertEquals(t, m, expected)
	})
}

func TestMaxChildIndex(t *testing.T) {
	t.Run("should return -1 if heap is empty", func(t *testing.T) {
		m := make(MinHeap, 0, 1)

		if c := m.MinChildIndex(0); c != -1 {
			t.Fatalf("expected -1 but got %v", c)
		}
	})

	t.Run("should return left child if right is empty", func(t *testing.T) {
		m := mockMinHeap([]float64{1, 2})

		if c := m.MinChildIndex(0); c != 1 {
			t.Fatalf("expected 1 but got %v", c)
		}
	})

	t.Run("should return right child if right child is smaller", func(t *testing.T) {
		m := mockMinHeap([]float64{1, 3, 2})

		if c := m.MinChildIndex(0); c != 2 {
			t.Fatalf("expected 2 but got %v", c)
		}
	})

	t.Run("should return right child if right child is smaller, child from the middle", func(t *testing.T) {
		m := mockMinHeap([]float64{1, 2, 3, 4, 5, 6, 7})

		if c := m.MinChildIndex(2); c != 5 {
			t.Fatalf("expected 2 but got %v", c)
		}
	})

	t.Run("should return left child if left child is smaller", func(t *testing.T) {
		m := mockMinHeap([]float64{1, 2, 3})

		if c := m.MinChildIndex(0); c != 1 {
			t.Fatalf("expected 2 but got %v", c)
		}
	})

	t.Run("should return left child if left child is smaller, from the middle", func(t *testing.T) {
		m := mockMinHeap([]float64{1, 2, 3, 4, 5, 7, 6})

		if c := m.MinChildIndex(2); c != 6 {
			t.Fatalf("expected 2 but got %v", c)
		}
	})

}

func TestPop(t *testing.T) {
	t.Run("should return nil if heap is empty", func(t *testing.T) {
		m := make(MinHeap, 0, 1)

		result := m.Pop()
		if result != nil {
			t.Fatalf("expected nil but got %v", result)
		}
	})

	t.Run("should empty the heap if there is only one element", func(t *testing.T) {
		m := mockMinHeap([]float64{1})

		result := m.Pop()
		if result.Value != 1 {
			t.Fatalf("expected nil but got %v", result)
		}

		if len(m) != 0 {
			t.Fatalf("expected 0 but got %v", len(m))
		}
	})

	t.Run("should return the root and bubble down", func(t *testing.T) {
		m := mockMinHeap([]float64{1, 2, 3, 4, 5, 6})

		root := m.Pop()

		if root.Value != 1 {
			t.Fatalf("expected 1 but got %v", root)
		}

		expected := []float64{2, 4, 3, 6, 5}
		assertEquals(t, m, expected)
	})

	t.Run("should return the root and bubble down (2)", func(t *testing.T) {
		m := mockMinHeap([]float64{1, 2, 3})

		root := m.Pop()

		if root.Value != 1 {
			t.Fatalf("expected 1 but got %v", root)
		}

		expected := []float64{2, 3}
		assertEquals(t, m, expected)
	})
}

func mockNode(value float64) *Node {
	return &Node{
		Value: value,
	}
}

func mockMinHeap(values []float64) MinHeap {
	m := make(MinHeap, 0, len(values))
	for _, v := range values {
		m = append(m, mockNode(v))
	}

	return m
}

func assertEquals(t *testing.T, m MinHeap, expected []float64) {
	var values []float64
	for _, v := range m {
		values = append(values, v.Value)
	}

	for i, v := range values {
		if v != expected[i] {
			t.Fatalf("expected %v but got %v", expected, values)
		}
	}
}
