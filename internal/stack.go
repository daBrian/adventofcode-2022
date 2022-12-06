package internal

type Stack[C any] struct {
	content []C
}

func (s *Stack[C]) IsEmpty() bool {
	return len(s.content) == 0
}

// Push a new value onto the Stack
func (s *Stack[C]) Push(str C) {
	s.content = append(s.content, str) // Simply append the new value to the end of the Stack
}

// Remove and return top element of Stack. Return false if Stack is empty.
func (s *Stack[C]) Pop() (C, bool) {
	if s.IsEmpty() {
		var noop C
		return noop, false
	} else {
		index := len(s.content) - 1     // Get the index of the top most element.
		element := (s.content)[index]   // Index into the slice and obtain the element.
		s.content = (s.content)[:index] // Remove it from the Stack by slicing it off.
		return element, true
	}
}

// Remove and return top element of Stack. Return false if Stack is empty.
func (s *Stack[C]) PopMore(n int) (rv []C, more bool) {
	var top C
	for i := 0; i < n; i++ {
		top, more = s.Pop()
		rv = append([]C{top}, rv...)
		if !more {
			break
		}
	}
	return
}

func (s *Stack[C]) Peek() (C, bool) {
	if s.IsEmpty() {
		var noop C
		return noop, false
	} else {
		index := len(s.content) - 1 // Get the index of the top most element.
		element := (s.content)[index]
		return element, true
	}
}

func (s *Stack[C]) PushMore(l []C) {
	for _, c := range l {
		s.Push(c)
	}
}
