package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"time"
)

const (
	//黑桃
	Spade = 0
	//红桃
	Hearts = 1
	//梅花
	Club = 2
	//方块
	Diamond = 3
)

type Poker struct {
	Num int
	Flower int
}

func (p Poker)PokerSelf()string  {
	var buffer string

	switch p.Flower {
	case Spade:
		buffer += "♤"
	case Hearts:
		buffer += "♡"
	case Club:
		buffer += "♧"
	case Diamond:
		buffer += "♢"
	}
	switch p.Num {
	case 13:
		buffer += "2"
	case 12:
		buffer += "A"
	case 11:
		buffer += "K"
	case 10:
		buffer += "Q"
	case 9:
		buffer += "J"
	default:
		buffer += strconv.Itoa(p.Num+2)
	}

	return buffer
}


func CreatePokers()(pokers Pokers)  {
	for i := 1; i < 14; i++ {
		for j := 0; j < 4; j++ {
			pokers = append(pokers,Poker{
				Num:    i,
				Flower: j,
			})
		}
	}
	return
}

type Pokers []Poker

//这里引用了 https://studygolang.com/articles/21260，加上的是自己的理解
//使用接口，保证了所有类型的切片都可以乱序，这个具有通用性，啥都能搅乱
func randSlice(slice interface{}) {
	//因为使用了接口，不知道对象类型，这里使用了反射的方法
	rv := reflect.ValueOf(slice) //获取值信息

	//不是切片的话，不能乱序，要返回
	if rv.Type().Kind() != reflect.Slice {
		return
	}

	//单身狗没人权
	length := rv.Len()
	if length < 2 {
		return
	}

	//按照切片类型类型返回交换方法
	swap := reflect.Swapper(slice)
	//随机种子由当前时间决定
	rand.Seed(time.Now().Unix())
	for i := length - 1; i >= 0; i-- {
		j := rand.Intn(length)   //rand.Intn生成随机数
		swap(i, j)  //快乐交换
	}
	return
}

func (p Pokers)Print()  {
	for _, i2 := range p {
		fmt.Print(i2.PokerSelf()," ")
	}
	fmt.Println()
}

//多值排序，为了满足接口的需要，给它加上素质三连函数
func (p Pokers) Len() int {
	return len(p)
}

func (p Pokers) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//自定义sort规则
func (p Pokers) Less(i, j int) bool {
	//首先是按照花色排序
	if p[i].Flower < p[j].Flower {
		return true
	} else if p[i].Flower == p[j].Flower {
		if p[i].Num < p[j].Num { //花色相同按值排序
			return true
		}
	}
	return false
}

func main() {
	pokers := CreatePokers()
	fmt.Println("洗牌前：")
	pokers.Print()
	randSlice(pokers)
	fmt.Println("洗牌后：")
	pokers.Print()
	sort.Sort(pokers)  //这个Sort其实也是用的接口，必须满足有素质三连函数才可以往里面放。
	fmt.Println("排序后：")
	pokers.Print()
}
