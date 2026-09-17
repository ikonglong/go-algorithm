package other

// 我自己写的版本
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	// 跟下面的版本相比，这里只滚动两个变量
	fn2, fn1 := 0, 1
	for i := 2; i <= n; i++ {
		fn := fn1 + fn2 // fn 临时变量
		fn2 = fn1       // fn2 滚动至 fn1
		fn1 = fn        // fn1 滚动至 fn
	}
	return fn1
}

// LeetCode 官方「滚动数组」版本
func fibRollingArray(n int) int {
	if n < 2 {
		return n
	}

	// 1e9 = 1 × 10⁹ = 1,000,000,000
	// const mod int = 1_000_000_007
	const mod int = 1e9 + 7

	// p 是 f(n-2)，初始时无法设置有实质意义的值，因为按照斐波那契数列的定义，F0 之前没有数列项了，只能设置为 0
	// q 是 f(n-1)，初始时是 f(0)
	// r 是 f(n)，初始时是 f(1)
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		// 初始时 p 的值不是 f(n-2)。
		// 为了计算新项，需要向前滚动。在向前滚动的过程中，p 就一直是 f(n-2) 了
		p = q // P 向前滚动至 Q
		q = r // Q 向前滚动至 R

		// 斐波那契数增长非常快，普通整数很快就无法保存了，所以题目规定取模
		// F(47) = 2,971,215,073       // 超过 32 位有符号整数上限
		// F(93) = 12,200,160,415,121,876,738 // 超过 64 位有符号整数上限
		r = (p + q) % mod // R 滚动至新产生的数列项
	}
	return r
}

// F(0) = 0，F(1) = 1
// F(n) = F(n - 1) + F(n - 2)，其中 n > 1
// 上面就是递推公式。只要写出递推公式，如何用递归实现就很清楚了

func fibRecurse(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibRecurse(n-1) + fibRecurse(n-2)
}
