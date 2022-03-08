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

// func toSlice(l *MagicList) []int {
// 	var itemList []int
// 	i := l.LastItem
// 	for true {
// 		if i != nil {
// 			itemList = append([]int{i.Value}, itemList...)
// 			i = i.PrevItem
// 		} else {
// 			break
// 		}
// 	}

// 	return itemList
// }

func toSlice(l *MagicList, itemList []int) []int {
	i := l.LastItem
	if i == nil {
		return itemList
	}

	itemList = append([]int{i.Value}, itemList...)
	l.LastItem = l.LastItem.PrevItem
	return toSlice(l, itemList)
}

func main() {
	magicList := &MagicList{}
	add(magicList, 111)
	add(magicList, 21211)
	add(magicList, 3)
	add(magicList, 4)
	add(magicList, 5)
	add(magicList, 7)
	add(magicList, 9)

	var itemList []int
	fmt.Println(toSlice(magicList, itemList))
}
