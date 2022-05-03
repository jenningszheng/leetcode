package divide_two_integer

const (
	maxNum = 2147483647
	minNum = -2147483648
)

func divide(a int, b int) int {
	absA := a
	absB := b
	ans := 0
	if a == 0 {
		return ans
	}
	if a < 0 {
		absA = -a
	}
	if b < 0 {
		absB = -b
	}
	da := absA
	db := absB
	if db == 1 {
		ans = da
	} else {
		// 减到不能减为止，除法就做完了
		// 不能减指的是相减后小于 0
		for da >= absB {
			db = absB
			k := 1
			// 在每次相减前，使用二分法找到尽可能大的被减数
			for da >= db+db {
				db += db
				k += k
			}
			da -= db
			ans += k
		}
	}
	if (a > 0 && b > 0) || (a < 0 && b < 0) {
		if ans > maxNum {
			return maxNum
		}
		return ans
	}
	if ans < minNum {
		return minNum
	}
	return -ans
}

func Divide(a int, b int) int {
	return divide(a, b)
}
