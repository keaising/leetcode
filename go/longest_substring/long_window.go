package long

type HashSet struct {
	hash map[string]bool
}

func (h *HashSet) exist(s string) bool {
	return h.hash[s]
}

func (h *HashSet) push(s string) {
	h.hash[s] = true
}

func (h *HashSet) remove(s string) {
	delete(h.hash, s)
}

func (h *HashSet) size() int {
	return len(h.hash)
}

func (h *HashSet) isEmpty() bool {
	return len(h.hash) == 0
}

func length(s string) int {
	var max = 0
	var set = &HashSet{map[string]bool{}}
	i, j := 0, 0
	for j < len(s) {
		if !set.exist(string(s[j])) || set.isEmpty() {
			set.push(string(s[j]))
			if set.size() > max {
				max = set.size()
			}
			j++
		} else {
			set.remove(string(s[i]))
			i++
		}
	}
	return max
}
