package shortest_distance_to_a_character

func ShortestToChar(s string, c byte) []int {
	n := len(s)
	pos := []int{n}
	ans := make([]int, n)
	for i := range s {
		if s[i] == c {
			pos = append(pos, i)
		}
	}
	pos = append(pos, n+n)
	j := 0
	l, r := pos[j], pos[j+1]
	ld, rd := 0, 0
	for i := range ans {
		if i > r {
			j++
			l = pos[j]
			r = pos[j+1]
		}
		ld = l - i
		rd = r - i
		if ld < 0 {
			ld *= -1
		}
		if rd < 0 {
			rd *= -1
		}
		if ld < rd {
			ans[i] = ld
		} else {
			ans[i] = rd
		}
	}
	return ans
}