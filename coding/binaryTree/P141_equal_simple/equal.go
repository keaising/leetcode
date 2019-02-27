package equal

import "../Data"

func equal(t1 *Data.Node, t2 *Data.Node) bool {
	return check(t1, t2) && equal(t1.Left, t2.Left) && equal(t1.Right, t2.Right)
}

func check (t1 *Data.Node, t2 *Data.Node) bool {
	if t2 == nil && t1 == nil {
		return true
	} else if (t1 != nil && t2 == nil) || (t1 == nil && t2 != nil) {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return check(t1.Left, t2.Left) && check(t1.Right, t2.Right)
}
