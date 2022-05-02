package main

import (
	"fmt"
	longest_absolute_file_path "leetcode/longest-absolute-file-path"
)

func main() {
	//s := "dir\n        file.txt"
	//s := "a\n\tb.txt\na2\n\tb2.txt"
	//s := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	//s := "file1.txt\nfile2.txt\nlongfile.txt"
	s := "dir\n\t        file.txt\n\tfile2.txt"
	//s := "a\n\taa\n\t\taaa\n\t\t\tfile1.txt\naaaaaaaaaaaaaaaaaaaaa\n\tsth.png"
	//s := "a\n\tb\n\t\tc\n\t\t\td\n\t\t\t\te.txt\n\t\t\t\talsdkjf.txt\n\t\tskdjfl.txtlsdkjflsdjflsajdflkjasklfjkasljfklas\n\tlskdjflkajsflj.txt"
	ans := longest_absolute_file_path.LengthLongestPath(s)
	fmt.Println(ans)
}
