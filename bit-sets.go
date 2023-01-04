package main

import (
	"fmt"

	"github.com/muscar/bit-sets/internal/bitset"
)

func main() {
	fmt.Println("opt")
	bs := bitset.New(8192)
	bs.Add(0)
	bs.Add(3)
	bs.Add(64)
	bs.Add(1023)
	bs.Add(1024)
	bs.Add(1025)
	bs.Add(8191)
	it1 := bs.Iter()
	for it1.Next() {
		fmt.Println(it1.Current())
	}

	fmt.Println("naive")
	bs = bitset.New(8192)
	bs.Add(0)
	bs.Add(3)
	bs.Add(64)
	bs.Add(1023)
	bs.Add(1024)
	bs.Add(1025)
	bs.Add(8191)
	it2 := bs.NaiveIter()
	for it2.Next() {
		fmt.Println(it2.Current())
	}

	fmt.Println("not great")
	bs = bitset.New(8192)
	bs.Add(0)
	bs.Add(3)
	bs.Add(64)
	bs.Add(1023)
	bs.Add(1024)
	bs.Add(1025)
	bs.Add(8191)
	it3 := bs.NotGreatIter()
	for it3.Next() {
		fmt.Println(it3.Current())
	}

	fmt.Println("not great #1")
	bs = bitset.New(8192)
	bs.Add(0)
	bs.Add(3)
	bs.Add(64)
	bs.Add(1023)
	bs.Add(1024)
	bs.Add(1025)
	bs.Add(8191)
	it4 := bs.NotGreatIter1()
	for it4.Next() {
		fmt.Println(it4.Current())
	}

	fmt.Println("naive chan")
	bs = bitset.New(8192)
	bs.Add(0)
	bs.Add(3)
	bs.Add(64)
	bs.Add(1023)
	bs.Add(1024)
	bs.Add(1025)
	bs.Add(8191)
	it5 := bs.ChanNaiveIter()
	for n := range it5 {
		fmt.Println(n)
	}

	fmt.Println("opt chan")
	bs = bitset.New(8192)
	bs.Add(0)
	bs.Add(3)
	bs.Add(64)
	bs.Add(1023)
	bs.Add(1024)
	bs.Add(1025)
	bs.Add(8191)
	it6 := bs.ChanOptIter()
	for n := range it6 {
		fmt.Println(n)
	}
}
