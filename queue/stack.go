package queue

import "errors"

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type LIFO struct {
	values []string
}

func (l *LIFO) Pop() (string, error) {
	if len(l.values) == 0 {
		return "", ErrEmptyStack
	}

	v := l.values[len(l.values)-1]
	l.values = l.values[:len(l.values)-1]
	return v, nil
}

func (l *LIFO) Append(s string) {
	l.values = append(l.values, s)
}

func (l *LIFO) Peek() (string, error) {
	if len(l.values) == 0 {
		return "", ErrEmptyStack
	}

	return l.values[len(l.values)-1], nil
}
