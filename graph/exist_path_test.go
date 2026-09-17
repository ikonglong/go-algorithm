package graph

import (
	"fmt"
	"testing"

	"github.com/ikonglong/go-algorithm/common"
	"github.com/stretchr/testify/assert"
)

// 这个题目并不关注每个节点的数据，只关注节点之间是否有路径连通，因此 edges 就是图的表示。
//
// 算法：
// 1. 创建一个队列，用来存储要处理的节点
// 2. 初始化队列，将起点添加到队列中
// 3. 只要队列不为空
//      取出队头元素，遍历其相邻节点
//        相邻节点是否是目标节点？
//        如果是，大功告成。如果不是，将此相邻节点的相邻节点添加到队列中
// 4. 返回 false，表明所有节点都遍历了，但没有找到
//

func validPath(n int, edges [][]int, source int, destination int) bool {
	if len(edges) == 0 || source == destination {
		return true
	}

	// 把图的表示从二维数组转换为映射表，因为二维表访问起来比较麻烦，
	// 要遍历一维数组才能找到一个节点的相邻节点
	graph := make(map[int][]int)
	for _, item := range edges {
		if neighbors, present := graph[item[0]]; present {
			graph[item[0]] = append(neighbors, item[1])
		} else {
			graph[item[0]] = []int{item[1]}
		}
		// 无向图，所以还需要添加反方向的映射
		if neighbors, present := graph[item[1]]; present {
			graph[item[1]] = append(neighbors, item[0])
		} else {
			graph[item[1]] = []int{item[0]}
		}
	}
	fmt.Printf("graph: %v\n", graph)

	que := make([]int, 0)

	// que = append(que, source)
	// for len(que) != 0 {
	// 	current := que[0]
	// 	que = que[1:] // 出队
	// 	for _, neighbor := range graph[current] {
	// 		if destination == neighbor {
	// 			return true
	// 		} else {
	// 			que = append(que, graph[neighbor]...)
	// 		}
	// 	}
	// }

	que = append(que, graph[source]...)
	searched := make(map[int]struct{})
	searched[source] = struct{}{}
	// 这个实现版本比上面注释掉的实现版本要简洁易懂。因为上面的版本
	// 在外层循环中搜索的不是当前节点，而是当前节点的相邻节点，这导致逻辑比较复杂
	for len(que) != 0 {
		current := que[0]
		que = que[1:]
		if _, present := searched[current]; present {
			continue
		}
		searched[current] = struct{}{}
		if destination == current {
			return true
		} else {
			que = append(que, graph[current]...)
		}
	}
	return false
}

func TestValidPath(t *testing.T) {
	exist := validPath(3,
		common.To2DArray("[[0,1],[1,2],[2,0]]"),
		0, 0)
	assert.Equal(t, true, exist)

	exist = validPath(10,
		common.To2DArray("[[0,7],[0,8],[6,1],[2,0],[0,4],[5,8],[4,7],[1,3],[3,5],[6,5]]"),
		7, 5)
	assert.Equal(t, true, exist)
}
