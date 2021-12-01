package main

func main() {

}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	idOrderMap := make(map[int]int)
	for i := range employees {
		idOrderMap[employees[i].Id] = i
	}

	var (
		queue []int
		res   int
	)
	queue = append(queue, id)
	for len(queue) > 0 {
		res += employees[idOrderMap[queue[0]]].Importance
		queue = append(queue, employees[idOrderMap[queue[0]]].Subordinates...)
		queue = queue[1:]
	}
	return res
}
