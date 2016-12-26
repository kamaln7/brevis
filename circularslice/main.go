package circularslice

type Slice struct {
	Size int
	data []interface{}
}

func New(size int) *Slice {
	return &Slice{
		Size: size,
	}
}

func (s *Slice) Insert(value interface{}) *Slice {
	s.data = append(s.data, value)
	if len(s.data) > s.Size {
		s.data = s.data[1:]
	}

	return s
}

func (s *Slice) Get() []interface{} {
	return s.data
}

func (s *Slice) Clear() {
	s.data = s.data[:0]
}
