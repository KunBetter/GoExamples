// LCS
// 最长公共子序列 The longest common subsequence
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(v1, v2 []int) int {
	/*
		设序列X=<x1, x2, …, xm>和Y=<y1, y2, …, yn>
		的一个最长公共子序列Z=<z1, z2, …, zk>，则：
		1. 若xm=yn，则zk=xm=yn 且Zk-1 是Xm-1 和Yn-1 的最长公共子序列；
		2. 若xm≠yn 且zk≠xm ，则Z 是Xm-1 和Y 的最长公共子序列；
		3. 若xm≠yn 且zk≠yn ，则Z 是X 和Yn-1 的最长公共子序列。
		其中Xm-1=<x1, x2, …, xm-1>，Yn-1=<y1, y2, …, yn-1>，Zk-1=<z1, z2, …, zk-1>。

		由最长公共子序列问题的最优子结构性质可知，要找出X=<x1, x2, …, xm>和Y=<y1, y2, …,
		yn>的最长公共子序列，可按以下方式递归地进行：当xm=yn 时，找出Xm-1 和Yn-1 的最长公
		共子序列，然后在其尾部加上xm(=yn)即可得X 和Y 的一个最长公共子序列。当xm≠yn 时，
		必须解两个子问题，即找出Xm-1 和Y 的一个最长公共子序列及X 和Yn-1 的一个最长公共子
		序列。这两个公共子序列中较长者即为X 和Y 的一个最长公共子序列。
		由此递归结构容易看到最长公共子序列问题具有子问题重叠性质。例如，在计算X 和Y
		的最长公共子序列时，可能要计算出X 和Yn-1 及Xm-1 和Y 的最长公共子序列。而这两个子
		问题都包含一个公共子问题，即计算Xm-1 和Yn-1 的最长公共子序列。
		与矩阵连乘积最优计算次序问题类似，我们来建立子问题的最优值的递归关系。用c[i,j]
		记录序列Xi 和Yj 的最长公共子序列的长度。其中Xi=<x1, x2, …, xi>，Yj=<y1, y2, …, yj>。当
		i=0 或j=0 时，空序列是Xi 和Yj 的最长公共子序列，故c[i,j]=0。其他情况下，由定理可建
		立递归关系如下：
				 0 							if i = 0 or j = 0
		c[i,j] = c[i-1,j-1]+1 				if i,j > 0 and xi = yj
				 max(c[i,j-1],c[i-1,j]) 	if i,j > 0 and xi != yj

		改进算法！
	*/
	v1Length := len(v1)
	v2Length := len(v2)

	if v1Length == 0 || v2Length == 0 {
		return 0
	}

	if v1[v1Length-1] == v2[v2Length-1] {
		return lcs(v1[:v1Length-1], v2[:v2Length-1]) + 1
	} else {
		return max(lcs(v1[:v1Length-1], v2[:v2Length]), lcs(v1[:v1Length], v2[:v2Length-1]))
	}
}

func main() {
	v1 := []int{1, 2, 3, 4, 5, 6}
	v2 := []int{0, 4, 5, 7}
	fmt.Println(lcs(v1, v2))
}
