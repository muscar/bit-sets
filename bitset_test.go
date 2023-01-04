package main

import (
	"testing"

	"github.com/muscar/bit-sets/internal/bitset"
)

func BenchmarkNaiveIterDense(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		bs.Add(i)
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.NaiveIter()
		for it.Next() {
			sum += it.Current()
		}
	}
}

func BenchmarkOptIterDense(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		bs.Add(i)
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.Iter()
		for it.Next() {
			sum += it.Current()
		}
	}
}

func BenchmarkChanNaiveIterDense(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		bs.Add(i)
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.ChanNaiveIter()
		for n := range it {
			sum += n
		}
	}
}

func BenchmarkChanOptIterDense(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		bs.Add(i)
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.ChanOptIter()
		for n := range it {
			sum += n
		}
	}
}

func BenchmarkNaiveIterSparse100(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%100 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.NaiveIter()
		for it.Next() {
			sum += it.Current()
		}
	}
}

func BenchmarkOptIterSparse100(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%100 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.Iter()
		for it.Next() {
			sum += it.Current()
		}
	}
}
func BenchmarkNaiveIterSparse1000(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%1000 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.NaiveIter()
		for it.Next() {
			sum += it.Current()
		}
	}
}

func BenchmarkOptIterSparse1000(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%1000 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.Iter()
		for it.Next() {
			sum += it.Current()
		}
	}
}
func BenchmarkNaiveIterSparse10000(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%10000 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.NaiveIter()
		for it.Next() {
			sum += it.Current()
		}
	}
}

func BenchmarkOptIterSparse10000(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%10000 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.Iter()
		for it.Next() {
			sum += it.Current()
		}
	}
}

func BenchmarkChanNaiveIterSparse100(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%100 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.ChanNaiveIter()
		for n := range it {
			sum += n
		}
	}
}

func BenchmarkChanOptIterSparse100(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%100 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.ChanOptIter()
		for n := range it {
			sum += n
		}
	}
}
func BenchmarkChanNaiveIterSparse1000(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%1000 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.ChanNaiveIter()
		for n := range it {
			sum += n
		}
	}
}

func BenchmarkChanOptIterSparse1000(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%1000 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.ChanOptIter()
		for n := range it {
			sum += n
		}
	}
}
func BenchmarkChanNaiveIterSparse10000(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%10000 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.ChanNaiveIter()
		for n := range it {
			sum += n
		}
	}
}

func BenchmarkChanOptIterSparse10000(b *testing.B) {
	bs := bitset.New(1000000)
	for i := 0; i < bs.Len(); i++ {
		if i%10000 == 0 {
			bs.Add(i)
		}
	}
	for i := 0; i < b.N; i++ {
		sum := 0
		it := bs.ChanOptIter()
		for n := range it {
			sum += n
		}
	}
}
