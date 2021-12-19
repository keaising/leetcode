package main

import (
	"container/heap"
	"log"
	"sort"
)

func main() {
	res := getOrder([][]int{
		{1, 2},
		{2, 4},
		{3, 2},
		{4, 1},
	})
	log.Println(res)
}

func getOrder(tasks [][]int) []int {
	var taskList []*Item
	for i, t := range tasks {
		taskList = append(taskList, &Item{
			Index:   i,
			Enqueue: t[0],
			Process: t[1],
		})
	}

	sort.Slice(taskList, func(i, j int) bool {
		return taskList[i].Enqueue < taskList[j].Enqueue
	})

	hp := Heap{}
	var res []int
	i, current := 0, 0
	for i < len(taskList) || hp.Len() > 0 {
		if hp.Len() == 0 && i < len(taskList) && taskList[i].Enqueue > current {
			current = taskList[i].Enqueue
		}

		for i < len(taskList) && taskList[i].Enqueue <= current {
			heap.Push(&hp, taskList[i])
			i++
		}

		if hp.Len() > 0 {
			top := heap.Pop(&hp).(*Item)
			current += top.Process
			res = append(res, top.Index)
		}
	}
	return res
}

type Item struct {
	Index   int
	Enqueue int
	Process int
}

type Heap []*Item

func (hp Heap) Len() int {
	return len(hp)
}

func (hp Heap) Less(i, j int) bool {
	if hp[i].Process == hp[j].Process {
		return hp[i].Index < hp[j].Index
	}
	return hp[i].Process < hp[j].Process
}

func (hp Heap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func (hp *Heap) Push(x interface{}) {
	item := x.(*Item)
	*hp = append(*hp, item)
}

func (hp *Heap) Pop() interface{} {
	n := len(*hp)
	item := (*hp)[n-1]
	(*hp)[n-1] = nil
	*hp = (*hp)[:n-1]
	return item
}
