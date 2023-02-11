package genericlist

type GenericList[T comparable] struct {
	data []T
}

func New[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (l *GenericList[T]) Insert(value T) {
	l.data = append(l.data, value)

}

func (l *GenericList[T]) Get(i int) T {
	if i > len(l.data)-1 {
		panic("Given Index Is Out Of Range")
	}

	for it := 0; it < len(l.data); it++ {
		if i == it {
			return l.data[it]
		}
	}

	panic("Value Not Found")
}

func (l *GenericList[T]) RemoveByValue(value T) {
	for i := 0; i < len(l.data); i++ {
		if l.data[i] == value {
			l.data = append(l.data[:i], l.data[i+1:]...)
		}
	}
}

func (l *GenericList[T]) Remove(i int) {
	if i > len(l.data)-1 {
		panic("Given Index Is Out Of Range")
	}

	for it := 0; it < len(l.data); it++ {
		if it == i {
			l.data = append(l.data[:it], l.data[it+1:]...)
		}
	}

}
