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
		case enter: // 中间节点已结束，此时要进入子节点或者返回到父节点
			i++
			nextLevel := 0
			// 计算下一节点所在的层级，以此来判断是进入子节点还是返回到父节点
			for i < n && input[i] == tab {
				i++
				nextLevel++
			}
			// 以下情况要返回到父节点：
			//   - nextLevel 为 0，表示下一节点是根节点的子节点
			//   - 下一层级比当前层级小
			//   - 字符串遍历完毕
			if nextLevel == 0 || nextLevel <= level || i >= n {
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
