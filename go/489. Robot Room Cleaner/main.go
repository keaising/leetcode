package main

import "fmt"

func main() {

}

/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 *
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */

type Robot struct {
}

func (robot *Robot) Move() bool {
	return false
}

func (robot *Robot) TurnLeft() {}

func (robot *Robot) TurnRight() {}

func (robot *Robot) Clean() {}

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func cleanRoom(robot *Robot) {
	var visit = make(map[string]int)
	robot.Clean()
	backtrack(robot, 0, 0, 0, visit)
}

func backtrack(robot *Robot, r, c, d int, visit map[string]int) {
	robot.Clean()
	visit[fmt.Sprintf("%d-%d", r, c)] = 0

	for range dirs {
		newR := r + dirs[d][0]
		newC := c + dirs[d][1]
		key := fmt.Sprintf("%d-%d", newR, newC)
		_, ok := visit[key]
		if !ok && robot.Move() {
			backtrack(robot, newR, newC, d, visit)
		}
		robot.TurnRight()
		d = (d + 1) % 4
	}
	goBack(robot)
}

func goBack(robot *Robot) {
	robot.TurnRight()
	robot.TurnRight()
	robot.Move()
	robot.TurnRight()
	robot.TurnRight()
}
