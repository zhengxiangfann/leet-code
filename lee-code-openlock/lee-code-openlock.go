package main

import (
	"fmt"
	"strconv"
)

// 1 <= deadends.length <= 500
// deadends[i].length == 4
// target.length == 4
// target 不在 deadends 之中
// target 和 deadends[i] 仅由若干位数字组成
func openLock(deadends []string, target string) int {
	// 将死亡数字放入一个集合中，方便快速判断是否为死亡状态
	deadSet := make(map[string]bool)
	for _, d := range deadends {
		deadSet[d] = true
	}

	// 如果初始状态就是死亡状态，则无法解锁，返回-1
	if deadSet["0000"] {
		return -1
	}

	// 使用队列来进行广度优先搜索
	queue := make([]string, 0)
	queue = append(queue, "0000")

	// 使用一个集合来记录已经访问过的状态，避免重复访问
	visited := make(map[string]bool)
	visited["0000"] = true

	// 记录转动次数
	count := 0

	// 开始广度优先搜索
	for len(queue) > 0 {
		size := len(queue)

		// 遍历当前层级的所有节点
		for i := 0; i < size; i++ {
			node := queue[i]

			// 如果当前节点是目标节点，返回转动次数
			if node == target {
				return count
			}

			// 将当前节点的相邻节点加入队列
			for _, next := range getNextStates(node) {
				// 如果相邻节点是死亡状态或已经访问过，则跳过
				if deadSet[next] || visited[next] {
					continue
				}

				// 将相邻节点加入队列和已访问集合
				queue = append(queue, next)
				visited[next] = true
			}
		}

		// 当前层级遍历完后，转动次数加一
		count++
		queue = queue[size:]
	}

	// 如果无法解锁，返回-1
	return -1
}

// 获取当前状态的相邻节点
func getNextStates(state string) []string {
	nextStates := make([]string, 0)

	for i := 0; i < len(state); i++ {
		// 获取当前位的数字
		digit := int(state[i] - '0')

		// 将数字向前旋转一位
		nextDigit1 := (digit + 1) % 10
		nextState1 := state[:i] + strconv.Itoa(nextDigit1) + state[i+1:]
		nextStates = append(nextStates, nextState1)

		// 将数字向后旋转一位
		nextDigit2 := (digit + 9) % 10
		nextState2 := state[:i] + strconv.Itoa(nextDigit2) + state[i+1:]
		nextStates = append(nextStates, nextState2)
	}

	return nextStates
}

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	ret := openLock(deadends, target)
	fmt.Println(ret)
}
