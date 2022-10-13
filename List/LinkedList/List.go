package main

import (
	"fmt"
)

func main() {
	list := List[int]{}

	list.Append(1, 2, 3, 4, 5, 6)
	fmt.Println(list.String())

	if err := list.DeleteByIndex(2); err != nil {
		fmt.Println(err)
	}

	list.Prepend(0)
	fmt.Println(list.String())

	if err := list.DeleteByItems(0, 1, 5); err != nil {
		fmt.Println(err)
	}
	fmt.Println(list.String())

	list.Append(10)

	v, err := list.GetByIndex(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}

type node[T any] struct {
	Next     *node[T]
	Previous *node[T]
	Data     T
}

type List[T any] struct {
	First *node[T]
	Last  *node[T]
	Count int
}

func (l *List[T]) Init(item T) {
	data := node[T]{}
	data.Data = item

	l.First = &data
	l.First.Previous = nil
	l.Last = &data
	l.Last.Next = nil

	l.Count = 1
}

func (l *List[T]) Append(items ...T) {
	if l.Last == nil {
		l.Init(items[0])

		items = items[1:]
	}

	for _, v := range items {
		data := node[T]{}
		data.Data = v

		l.Last.Next = &data
		prev := l.Last
		l.Last = &data
		l.Last.Previous = prev

		l.Count++
	}

}

func (l *List[T]) Prepend(items ...T) {
	for _, v := range items {
		data := node[T]{}
		data.Data = v

		l.First.Previous = &data
		next := l.First
		l.First = &data
		l.First.Next = next

		l.Count++
	}
}

func (l *List[T]) DeleteByIndex(index int) error {
	if errIndex := l.outIndex(index); errIndex != nil {
		return errIndex
	}

	item := l.First
	for i := 0; i < index; i++ {
		item = item.Next
	}

	l.delete(item)

	return nil
}

func (l *List[T]) DeleteByItems(items ...T) (err error) {
	if len(items) > l.Count {
		return fmt.Errorf("index out of range")
	}

	mapa := l.convertToMap(items)
	for item := l.First; item.Next != nil; {
		_, ok := mapa[item.Data]
		if ok {
			prev := item.Previous
			if prev == nil {
				next := item.Next
				l.delete(item)
				item = next
				l.First = item
				continue
			}

			l.delete(item)
			item = prev
			item = item.Next
			continue
		}
		item = item.Next
	}

	return nil
}

func (l *List[T]) delete(item *node[T]) {
	if item.Previous != nil {
		item.Previous.Next = item.Next
	} else {
		l.First = item.Next
	}

	if item.Next != nil {
		item.Next.Previous = item.Previous
	}

	item.Next = nil
	item.Previous = nil
	l.Count--
}

func (l *List[T]) GetByIndex(index int) (T, error) {
	if err := l.outIndex(index); err != nil {
		return *new(T), err
	}

	item := l.First
	for i := 0; i < index; i++ {
		item = item.Next
	}

	return item.Data, nil
}

func (l *List[T]) String() string {
	var str = "["
	for data := l.First; data != nil; data = data.Next {
		str += fmt.Sprintf("%v ", data.Data)
	}
	str = str[0:len(str)-1] + "]"

	return str
}

func (l *List[T]) convertToMap(items []T) (mapa map[any]struct{}) {
	mapa = make(map[any]struct{})
	for _, v := range items {
		mapa[v] = struct{}{}
	}

	return mapa
}

func (l *List[T]) outIndex(index int) error {
	if index >= l.Count {
		return fmt.Errorf("index out of range! Count: %v, Index: %v", l.Count, index)
	}

	return nil
}
