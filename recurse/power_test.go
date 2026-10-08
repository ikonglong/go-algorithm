package recurse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
// n 可能为负数

// 快速幂（分治） + 递归
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / recursePow(x, -n)
	}
	return recursePow(x, n)
}

// x^77 -> (x^38)^2 * x, x^38 -> (x^19)^2, x^19 -> (x^9)^2 * x,
// x^9 -> (x^4)^2 * x, x^4 -> (x^2)^2, x^2 -> (x^1)^2,
// x^1 -> (x^0)^2 * x
func recursePow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	subR := recursePow(x, n/2)
	r := subR * subR
	if n&1 == 1 { // n 是奇数
		return r * x
	}
	return r
}

// 快速幂（分治） + 迭代/循环
//
// 结合实例分析推导：x^77 => x -> x^2 -> x^4 -> x^8*x -> x^18*x -> x^38 -> x^76*x
// x^18*x 中最后的 x 此后被平方了 1 次，因此对 x^77 的贡献是 x^1
// x^18*x 中最后的 x 此后被平方了 2 次，因此对 x^77 的贡献是 (x^2)^2 = x^4
// x^8*x 中最后的 x 此后被平方了 3 次，因此对 x^77 的贡献是 (x^2)^3 = x^8
// 最初的 x 此后被平方了 6 次，因此对 x^77 的贡献是 (x^2)^6 = x^64
// 77 的二进制表示：1001101
// 1, 4, 8, 64 和刚好为 77，因为它们分别对应 77 二进制表示中值为 1 的位的十进制表示
//
// n 的二进制（从高到低）表示：2^ik + ... + 2^i1 + 2^i0
// x^n = x^(2^ik) + ... + x^(2^i1) + x^(2^i0)
// 这是发现的规律。数学原理，还得想想

func myPow2(x float64, n int) float64 {
	// TODO
	return float64(0)
}

// 采用递归实现：
//  1. x^n 跟 x^n-1 解决思路一样
//  2. 递推公式：x^n = x * x^n-1
//     终止条件：n-1>=0 => n>=1, x^0=1
//
// 该实现版本执行时报告：超出内存限制
func myPow3(x float64, n int) float64 {
	if n < 0 {
		return 1 / recursePow3(x, -n)
	}
	return recursePow3(x, n)
}

func recursePow3(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	return x * recursePow3(x, n-1)
}

func TestMyPow(t *testing.T) {
	assert.Equal(t, float64(1), myPow3(2, 0))
	assert.Equal(t, float64(8), myPow3(2, 3))
	assert.Equal(t, float64(1024), myPow3(2, 10))
	assert.Equal(t, float64(100), myPow3(10, 2))
}
