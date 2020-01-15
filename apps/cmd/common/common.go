package main

import "fmt"

const (
	x = iota * 10
	y
	z = "yy"
	k
	_
	p = iota
)

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

//作用域和函数执行
func doDefer() {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func doAppend() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func doIota() {
	fmt.Println(x, y, z, k, p)
}

//常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用。
func doConst() {
	const cl = 100

	var bl = 123
	println(&bl, bl)
	//println(&cl,cl)
}

//底层类型为int类型，但是不能直接赋值，需要强转
func doType() {
	type MyInt1 int
	//基于一个类型创建一个新类型，称之为defintion；基于一个类型创建一个别名，称之为alias
	type MyInt2 = int
	var i int = 9
	//var i1 MyInt1 = i
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

func doPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}

func main() {
	//:defer和函数返回值
	doDefer()
	//append append切片时候别漏了’…’
	doAppend()
	doIota()
	doPanic()
}
