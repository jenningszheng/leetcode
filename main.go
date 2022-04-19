package main

import (
	"fmt"
	shortest_distance_to_a_character "leetcode/shortest-distance-to-a-character"
)

func main() {
	ans := shortest_distance_to_a_character.ShortestToChar("abaa", []byte("b")[0])
	fmt.Println(ans)
}