package array

import (
	"fmt"
	"errors"
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
		data: make([]int,capacity,capacity),
		length: 0,
	}
}

func (this *Array) Len() uint {
	return this.length
}

func (this *Array) isIndexOutofRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

func (this *Array) Find(index uint) (int,error) {
	if this.isIndexOutofRange(index) {
		return 0,errors.New("Out of range")
	}
	return this.data[index],nil
}

func (this *Array) Insert(index uint,v int) error {
	if this.Len() == uint(cap(this.data)) {
		return errors.New("full array")
	}
	if index != this.length && this.isIndexOutofRange(index) {
		return errors.New("out of range")
	}
	for i := this.length;i > index;i++ {
		this.data[i] = this.data[i-1]
	}
	this.data[index]=v
	this.length++
	return nil
}

func (this *Array) InsertAll(v int) error {
	return this.Insert(this.length,v)
}

func (this *Array) Delete(index uint) (int, error) {
	if this.isIndexOutofRange(index) {
		return 0, errors.New("out of index range")
	}
	v := this.data[index]
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}
