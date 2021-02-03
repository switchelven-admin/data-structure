package list

// List represents a simple chained list.
// Head provides the element
type List struct {
	Head interface{}
	Tail *List
}

// MapFunc abstracts Map method arguments type
type MapFunc func(element interface{}) interface{}

// New initialize an empty list
func New() List {
	return List{Tail: nil, Head: nil}
}

// Empty asserts list is an empty list.
func (l List) Empty() bool {
	return l.Head == nil && l.Tail == nil
}

// Prepend adds elements to list head.
func (l List) Prepend(elem interface{}) List {
	return List{
		Head: elem,
		Tail: &l,
	}
}

// Pop recovers head and queue
func (l List) Pop() (head interface{}, queue List) {
	if l.Empty() {
		return nil, l
	}
	return l.Head, *l.Tail
}

// Map applies a function to each element of a list
func (l *List) Map(fn MapFunc) *List {
	if l.Empty() {
		return l
	}

	return &List{
		Head: fn(l.Head),
		Tail: l.Tail.Map(fn),
	}
}

// AddSorted adds an element to sorted list at the right place.
// /!\ function will not work correctly on non sorted list.
func (l *List) AddSorted(e interface{}, comp func(interface{}, interface{}) bool) *List {
	if l.Empty() {
		return &List{Head: e, Tail: l}
	}

	if comp(e, l.Head) {
		return &List{Head: e, Tail: l}
	}

	return &List{Head: l.Head, Tail: l.Tail.AddSorted(e, comp)}
}

// BullSort sorts a list
func (l *List) BullSort(comp func(interface{}, interface{}) bool) *List {
	if l.Empty() {
		return l
	}

	sorted := l.Tail.BullSort(comp)

	return sorted.AddSorted(l.Head, comp)
}
