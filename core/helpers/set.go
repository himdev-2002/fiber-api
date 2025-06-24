package helpers

// Generic Set
type Set[T comparable] struct {
	m map[T]struct{}
}

// Buat set baru
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

// Tambah elemen ke set
func (s *Set[T]) Add(val T) {
	s.m[val] = struct{}{}
}

// Hapus elemen dari set
func (s *Set[T]) Remove(val T) {
	delete(s.m, val)
}

// Cek apakah elemen ada
func (s *Set[T]) Contains(val T) bool {
	_, ok := s.m[val]
	return ok
}

// Ambil semua elemen sebagai slice
func (s *Set[T]) Values() []T {
	vals := make([]T, 0, len(s.m))
	for k := range s.m {
		vals = append(vals, k)
	}
	return vals
}

// Union: gabungkan dua set
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for k := range s.m {
		result.Add(k)
	}
	for k := range other.m {
		result.Add(k)
	}
	return result
}

// Intersection: elemen yang sama
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for k := range s.m {
		if other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

// Difference: elemen di s tapi tidak di other
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for k := range s.m {
		if !other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}
