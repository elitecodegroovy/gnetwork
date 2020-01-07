package main

//单例模式测试代码
func singleton() {
	counter1 := GetInstance()
	counter1.AddOne()
	println("count 1 : ", counter1.count)

	counter2 := GetInstance()
	counter2.AddOne()
	println("count 2 : ", counter1.count)
}

func main() {
	singleton()

}
