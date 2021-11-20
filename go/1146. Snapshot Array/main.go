package main

import (
	"fmt"
	"log"
)

func main() {
	obj := Constructor(3)
	obj.Set(1, 18)
	obj.Set(1, 4)
	s := obj.Snap()
	fmt.Println("snap: ", s)

	g := obj.Get(0, 0)
	fmt.Println("get: ", g)

	obj.Set(0, 20)
	s = obj.Snap()
	fmt.Println("snap: ", s)

	obj.Set(0, 2)
	obj.Set(1, 1)
	fmt.Println("set 1, 1: ", s, obj.m)
	// fmt.Printf("%#v\n", obj.m)

	g = obj.Get(1, 1)
	fmt.Println("get 1, 1: ", g)
	g = obj.Get(1, 0)
	fmt.Println("get: ", g)
}

type SnapshotArray struct {
	snapID int
	m      map[int][]snap
}

type snap struct {
	snapID int
	v      int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		snapID: 0,
		m:      make(map[int][]snap),
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	if arr, ok := this.m[index]; ok {
		if arr[len(arr)-1].snapID == this.snapID {
			this.m[index][len(arr)-1].v = val
			return
		}
	}
	this.m[index] = append(this.m[index], snap{
		snapID: this.snapID,
		v:      val,
	})
}

func (this *SnapshotArray) Snap() int {
	this.snapID += 1
	return this.snapID - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	arr, ok := this.m[index]
	if !ok {
		return 0
	}
	if arr[0].snapID > snap_id {
		return 0
	}
	if arr[len(arr)-1].snapID <= snap_id {
		return arr[len(arr)-1].v
	}

	var (
		l = 0
		r = len(arr) - 1
	)
	if index == 1 && snap_id == 1 {
		log.Println(arr, l, r, "snap_id", snap_id)
	}

	for l <= r {
		m := l + (r-l)/2
		if index == 1 && snap_id == 1 {
			log.Println("loop 1", m, arr[m])
			log.Println("loop 2", arr[m].snapID, snap_id)
		}

		if arr[m].snapID > snap_id {
			r = m - 1
		} else if arr[m].snapID < snap_id {
			l = m + 1
		} else {
			return arr[m].v
		}
		if index == 1 && snap_id == 1 {
			log.Println("loop 3", l, r)
		}
	}
	return arr[l].v
}
