package structures

import (
	"testing"
)

type MockQueue struct {
	data []any
}

func (m *MockQueue) Push(v any) {
	m.data = append(m.data, v)
}

func (m *MockQueue) Pop() (any, bool) {
	if len(m.data) == 0 {
		return nil, false
	}
	v := m.data
	m.data = m.data[1:]
	return v, true
}

func TestQueueInterface(t *testing.T) {
	var q Queue = &MockQueue{}
	q.Push(1)
	q.Push(2)

	v, ok := q.Pop()
	if !ok || v != 1 {
		t.Errorf("Pop() вернул неверный результат: %v, %v", v, ok)
	}

	v, ok = q.Pop()
	if !ok || v != 2 {
		t.Errorf("Pop() вернул неверный результат: %v, %v", v, ok)
	}

	_, ok = q.Pop()
	if ok {
		t.Error("Pop() должен вернуть false при пустой очереди")
	}
}

type MockHeap struct {
	data []any
}

func (m *MockHeap) Push(v any) {
	m.data = append(m.data, v)
}

func (m *MockHeap) Pop() (any, bool) {
	if len(m.data) == 0 {
		return nil, false
	}
	v := m.data
	m.data = m.data[1:]
	return v, true
}

func TestHeapInterface(t *testing.T) {
	var h Heap = &MockHeap{}
	h.Push(10)
	h.Push(20)

	v, ok := h.Pop()
	if !ok {
		t.Error("Pop() должен вернуть true при непустой куче")
	}
	if v == nil {
		t.Error("Pop() вернул nil вместо значения")
	}

	_, ok = h.Pop()
	if !ok {
		t.Error("Pop() должен вернуть true при непустой куче")
	}

	_, ok = h.Pop()
	if ok {
		t.Error("Pop() должен вернуть false при пустой куче")
	}
}

type MockBinaryTree struct {
	data []any
}

func (m *MockBinaryTree) Push(v any) {
	m.data = append(m.data, v)
}

func (m *MockBinaryTree) Pop() (any, bool) {
	if len(m.data) == 0 {
		return nil, false
	}
	v := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	return v, true
}

func (m *MockBinaryTree) Search(v any) bool {
	for _, item := range m.data {
		if item == v {
			return true
		}
	}
	return false
}

func TestBinaryTreeInterface(t *testing.T) {
	var bt BinaryTree = &MockBinaryTree{}
	bt.Push(5)
	bt.Push(10)
	bt.Push(15)

	if !bt.Search(10) {
		t.Error("Search(10) должен вернуть true")
	}

	if bt.Search(99) {
		t.Error("Search(99) должен вернуть false")
	}

	v, ok := bt.Pop()
	if !ok || v == nil {
		t.Error("Pop() должен вернуть значение и true")
	}

	_, ok = bt.Pop()
	if !ok {
		t.Error("Pop() должен вернуть true при непустом дереве")
	}

	_, ok = bt.Pop()
	if !ok {
		t.Error("Pop() должен вернуть true при непустом дереве")
	}

	_, ok = bt.Pop()
	if ok {
		t.Error("Pop() должен вернуть false при пустом дереве")
	}
}
