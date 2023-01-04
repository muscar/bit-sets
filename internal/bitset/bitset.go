package bitset

type Iterator interface {
	Next() bool
	Current() int
}

type NaiveIter struct {
	s   *BitSet
	off int
}

type NotGreatIter struct {
	s          *BitSet
	word       uint64
	woff, boff int
}

type NotGreatIter1 struct {
	s          *BitSet
	word       uint64
	woff, boff int
}

type Iter struct {
	s          *BitSet
	word       uint64
	woff, boff int
}

type BitSet struct {
	words    []uint64
	wordsLen int
	len      int
}

func New(size int) *BitSet {
	words := size / 64
	// add an extra word as a sentinel
	return &BitSet{make([]uint64, words+1), words, words * 64}
}

func (s *BitSet) Len() int {
	return s.len
}

func (s *BitSet) Add(n int) {
	word, bit := n/64, n%64
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *BitSet) Has(n int) bool {
	word, bit := n/64, n%64
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *BitSet) NaiveIter() *NaiveIter {
	return &NaiveIter{s, -1}
}

func (s *BitSet) NotGreatIter() *NotGreatIter {
	return &NotGreatIter{s, s.words[0], 0, -1}
}

func (s *BitSet) NotGreatIter1() *NotGreatIter1 {
	return &NotGreatIter1{s, s.words[0], 0, -1}
}

func (s *BitSet) Iter() *Iter {
	return &Iter{s, s.words[0], 0, -1}
}

func (s *BitSet) ChanNaiveIter() chan int {
	ch := make(chan int, 1)
	go func() {
		for i := 0; i < s.len; i++ {
			if s.Has(i) {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func (s *BitSet) ChanOptIter() chan int {
	ch := make(chan int, 1)
	go func() {
		for word := 0; word < s.wordsLen; word++ {
			segment := s.words[word]
			for bit := 0; segment > 0; bit++ {
				if (segment & 0x01) != 0 {
					ch <- word*64 + bit
				}
				segment >>= 1
			}
		}
		close(ch)
	}()
	return ch
}

// Naive iter

func (it *NaiveIter) Next() bool {
	it.off++
	for it.off < it.s.len && !it.s.Has(it.off) {
		it.off++
	}
	return it.off < it.s.len
}

func (it *NaiveIter) Current() int {
	return it.off
}

// Not great iter

func (it *NotGreatIter) Next() bool {
	if it.woff >= len(it.s.words) {
		return false
	}
	// ensure we make progress
	it.boff++
	// find the first set bit in this word
	for it.boff < 64 && it.s.words[it.woff]&(1<<it.boff) == 0 {
		it.boff++
	}
	// we found a set bit
	if it.boff < 64 {
		return true
	}
	// we need to find the next non-zero word
	it.boff = 0
	it.woff++
	for it.woff < len(it.s.words) && it.s.words[it.woff] == 0 {
		it.woff++
	}
	// we're done with the whole set
	if it.woff >= len(it.s.words) {
		return false
	}
	// we found a non-zero word; find the first set bit
	for it.s.words[it.woff]&(1<<it.boff) == 0 {
		it.boff++
	}
	return true
}

func (it *NotGreatIter) Current() int {
	return it.woff*64 + it.boff
}

// Not great iter #1

func (it *NotGreatIter1) Next() bool {
	for it.woff < it.s.wordsLen {
		// ensure we make progress
		it.boff++
		// find the first set bit in this word
		for it.word != 0 && it.word&0x01 == 0 {
			it.boff++
			it.word >>= 1
		}
		// we found a set bit
		if it.word > 0 {
			it.word >>= 1
			return true
		}
		// we need to find the next non-zero word
		it.boff = -1
		it.woff++
		it.word = it.s.words[it.woff]
	}
	return false
}

func (it *NotGreatIter1) Current() int {
	return it.woff*64 + it.boff
}

// Iter

func (it *Iter) Next() bool {
	for it.woff < it.s.wordsLen {
		for it.boff++; it.word > 0 && it.word&0x01 == 0; it.word >>= 1 {
			it.boff++
		}
		if it.word > 0 {
			it.word >>= 1
			return true
		}
		it.woff++
		// OK, we have a sentinel
		it.word = it.s.words[it.woff]
		it.boff = -1
	}
	return false
}

func (it *Iter) Current() int {
	return it.woff*64 + it.boff
}
