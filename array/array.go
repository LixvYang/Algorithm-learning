package array

import (
	"errors"
	"fmt"
)

type Array struct {
	data []int
	length uint
}

func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data: make([]int,capacity),
		length: 0,
	}
}

func (a *Array) Len() uint {
	return a.length
}

func (a *Array) isIndexOutOfRange(index uint) bool {
	if index > uint(cap(a.data)) {
		return true
	} else {
		return false
	}
}

func (a *Array) Find(index uint) (int,error) {
	if a.isIndexOutOfRange(index) {
		return 0, fmt.Errorf("out of index range")
	}
	return a.data[index],nil
}

func (a *Array) Insert(index uint,v int) error {
	if a.Len() == uint(cap(a.data)) {
		return errors.New("full array")
	}
	if index != a.length && a.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	for i:= a.length;i > index;i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.length++
	return nil
}

func (a *Array) InsertToTail(v int) error {
	return a.Insert(a.Len(), v)
}

func (a *Array) Delete(index uint) (int,error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := a.data[index]
	for i := index;i < a.Len() - 1;i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return v,nil
}

func (a *Array) Print()  {
	var format string

	for i := uint(0); i < a.Len(); i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}