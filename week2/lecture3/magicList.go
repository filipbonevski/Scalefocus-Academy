package main

import "fmt"

type MagicList struct {
	LastItem *Item
}

type Item struct {
	Value    int
	PrevItem *Item
}

func add(l *MagicList, value int) {
	i := Item{Value: value}

	if l.LastItem == nil {
		l.LastItem = &i
	} else {
		i.PrevItem = l.LastItem
		l.LastItem = &i
	}

}

func toSlice(l *MagicList) []int {
	var itemList []int
	i := l.LastItem
	for true {
		if i != nil {
			itemList = append([]int{i.Value}, itemList...)
			i = i.PrevItem
		} else {
			break
		}
	}

	return itemList
}

func main() {
	magicList := &MagicList{}
	add(magicList, 111)
	add(magicList, 21211)
	add(magicList, 3)
	add(magicList, 4)
	add(magicList, 5)

	fmt.Println(toSlice(magicList))
}
