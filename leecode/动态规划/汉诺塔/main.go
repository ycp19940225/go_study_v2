package main

import "fmt"

func main() {

	// 左到右有A、B、C三根柱子，其中A柱子上面有从小叠到大的n个圆盘，
	// 现要求将A柱子上的圆盘移到C柱子上去，期间只有一个原则：
	// 一次只能移到一个盘子且大盘子不能在小盘子上面，求移动的步骤和移动的次数

	// 1个 1 -> c
	// 2个 1 -> b 2 -> c 1 -> c  || 1 -> c 2->b 1 -> a 2 -> c 1 -> c
	// 3个 1 -> c 2-> b 1 -> b 3->c 1 -> a 2 -> c 1->c
	// 4个 1 -> c 2-> b 1 -> b 3->c 1 -> a 2 -> c 1->c
	// f(n) = f(n-1)+1
	// f(3) = 7
	// f(2) = 5
	// f(1) = 1

	//res := test(5)
	//
	//println(res)
	//hanoid(4, "a", "b", "c")

	// 细胞分裂 有一个细胞 每一个小时分裂一次，一次分裂一个子细胞，第三个小时后会死亡。那么n个小时候有多少细胞？
	// 0 1个细胞 a
	// 1 2个细胞 a b
	// 2 4个细胞 a b c d
	// 3 8个细胞 a b c d e f g h - 死亡的细胞 1 a
	// 3 18个细胞 a b c d e f g h - 死亡的细胞 1 a
	// f(n) = f(n-1-f(n-3)) + f(n-2) + f(n-3)
}

// f(n, a, b, c) = f(n-1,a, c, b) + move(a, c) + f(n-1, b, a, c)
func hanoid(n int, a, b, c string) int {
	if n == 0 {
		return 0
	}
	hanoid(n-1, a, c, b)
	move(n-1, a, c)
	hanoid(n-1, b, a, c)
	return 0
}

func move(n int, _ string, c string) {
	fmt.Printf("%d: -> %s \n", n+1, c)
}

func test(n int) int {
	if n == 3 {
		return 7
	}
	if n == 2 {
		return 1
	}
	if n == 1 {
		return 1
	}

	return test(n-1) + 1
}
