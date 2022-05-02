package lexicographical_numbers

func LexicalOrder(n int) []int {
	ans := make([]int, n)
	num := 1
	for i := range ans {
		ans[i] = num
		if num*10 <= n {
			num = num * 10
		} else {
			for {
				if num%10 == 9 || num+1 > n {
					num = num / 10
					continue
				}
				break
			}
			num++
		}
	}
	return ans
}
