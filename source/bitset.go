package main

type BitSet uint

func Bit(i int) uint {
	return 1 << uint(i)
}

func (s *BitSet) Add(i int) {
	*s |= BitSet(Bit(i))
}

func (s *BitSet) Remove(i int) {
	*s &= BitSet(^Bit(i))
}

func (s BitSet) Contains(i int) bool {
	return s&BitSet(Bit(i)) > 0
}

func (s BitSet) IsFull(n int) bool {
	return s == BitSet(Bit(n)-1)
}
