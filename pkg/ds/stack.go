package ds

type node struct {
	item string
	next *node
}

type StackLinkedListImpl struct {
	first *node
	size  int
}

func (s *StackLinkedListImpl) GetSize() int {
	return s.size
}

func (s *StackLinkedListImpl) IsEmpty() bool {
	return s.size == 0
}

func (s *StackLinkedListImpl) Push(v string) {
	oldFirst := s.first
	s.first = &node{
		item: v,
		next: oldFirst,
	}

	s.size++
}

func (s *StackLinkedListImpl) Pop() string {
	oldFirst := s.first
	s.first = s.first.next

	s.size--
	return oldFirst.item
}
