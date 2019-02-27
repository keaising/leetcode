package contains

import "../Data"

func contains (t1 *Data.Node, t2 *Data.Node) bool {
	return contains(t1, t2) || check(t1, t2)
}

func check (t1 *Data.Node, t2 *Data.Node) bool {
	if t2 == nil  {
		return true
	}
	if t1 == nil || t1.Val != t2.Val {
		return false
	}
	return check(t1.Left, t2.Left) && check(t1.Right, t2.Right)
}
