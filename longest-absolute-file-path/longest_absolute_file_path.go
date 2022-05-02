package longest_absolute_file_path

const (
	tab   = 9
	enter = 10
	dot   = 46
)

func longestPath(input string, level int) (int, int) {
	l := 0
	ans := 0
	i := 0
	n := len(input)
	for i < n {
		ch := input[i]
		switch ch {
		case enter: // 中间节点
			i++
			nextLevel := 0
			// 遇到 \t 的时候，表示要进入子节点或者返回到父节点
			for i < n && input[i] == tab {
				i++
				nextLevel++
			}
			// 层级为0，直接返回到根节点
			if nextLevel == 0 {
				return ans, i - 1
			}
			// 下一层级比当前层级小，要返回到父节点
			if nextLevel <= level || i >= n {
				return ans, i - nextLevel - 1
			}
			// 进入子节点
			sl, ss := longestPath(input[i:], nextLevel)
			if sl > 0 && l+1+sl > ans {
				ans = l + 1 + sl
			}
			i += ss
		case dot: // 叶子节点
			i++
			l++
			for i < n && input[i] != enter {
				i++
				l++
			}
			return l, i
		default:
			l++
			i++
		}
	}
	return ans, i
}

func LengthLongestPath(input string) int {
	n := len(input)
	ans := 0
	i := 0
	for i < n {
		// 根路径前面的 \n 是没有意义的，要去掉
		for i < n && input[i] == enter {
			i++
		}
		s := input[i:]
		l, ls := longestPath(s, 0)
		i += ls
		if l > ans {
			ans = l
		}
	}
	return ans
}
