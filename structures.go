package structures

// Queue определяет интерфейс очереди с операциями push и pop.
type Queue interface {
	Push(v any)
	Pop() (any, bool)
}

// Heap определяет интерфейс кучи с операциями push и pop.
type Heap interface {
	Push(v any)
	Pop() (any, bool)
}

// BinaryTree определяет интерфейс бинарного дерева с операциями push, pop и search.
type BinaryTree interface {
	Push(v any)
	Pop() (any, bool)
	Search(v any) bool
}
