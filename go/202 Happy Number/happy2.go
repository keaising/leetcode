package main

func isHappy2(n int) bool {
	if n == 0 {
		return false
	}
	happy := func(i int) int {
		ans := 0
		for i > 0 {
			ans += (i % 10) * (i % 10)
			i = i / 10
		}
		return ans
	}

	walker, runner := n, n
	for {
		walker = happy(n)
		runner = happy(happy(n))

		if walker == runner {
			break
		}
	}
	return walker == 1
}
