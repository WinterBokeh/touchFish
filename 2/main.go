package main

import (
	"fmt"
	"sync"
)

//卡时间
var wA = sync.WaitGroup{}
var wB = sync.WaitGroup{}

//优化：
//这里使用约数和定理，代码是之前oi时代的板子，改写成go直接用上了
func sum(n int) int {
	s := 1
	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			a := 1
			for {
				if n % i != 0 {
					break
				}
				n /= i
				a *= i
			}
			s = s * ( a * i - 1 ) / ( i - 1 )
		}
	}
	if n > 1 {
		s = s * (n + 1)
	}
	return s
}

//求素数
func getPrimeNum() {
	for i := 4; i <= 123456; i++ {
		wB.Add(1)
		go goJudgePrime(i)
	}
}

//求完全数
func getPerfectNum()  {
	for i := 6; i <= 123456; i++ {
		wA.Add(1)
		go goJudgePerfect(i)
	}
}

//判断是否为完全数
func goJudgePerfect(n int) {
	if n == sum(n) - n {
		fmt.Print(n, " ")
	}

	wA.Done()
}

//判断是否为素数
func goJudgePrime(n int) {
	flag := true

	for i := 2; i <= n-1; i++ {
		if (n % i) == 0 {
			flag = false
			break
		}
	}

	if flag {
		fmt.Print(n, " ")
	}

	wB.Done()
}

func main() {
	getPerfectNum()
	fmt.Println()
	getPrimeNum()
}
