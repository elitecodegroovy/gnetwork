package main

import (
	"bufio"
	"fmt"
	workpool "github.com/elitecodegroovy/gnetwork/apps/pool"
	"os"
	"runtime"
	"strconv"
	"time"
)

type MyWork struct {
	Name      string "The Name of a person"
	BirthYear int    "The Yea the person was born"
	WP        *workpool.WorkPool
}

// 线程执行单元
func (workPool *MyWork) DoWork(workRoutine int) {
	fmt.Printf("任务名称： %s : %d\n", workPool.Name, workPool.BirthYear)
	fmt.Printf(">>> workRoutine: %d，  QueuedWork: %d ， ActiveRoutines: %d\n", workRoutine, workPool.WP.QueuedWork(), workPool.WP.ActiveRoutines())
	time.Sleep(100 * time.Millisecond)

	//panic("test")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	workerWorkPool := workpool.New(runtime.NumCPU()*3, 100)
	shutdown := false

	go func() {
		for i := 0; i < 1000; i++ {
			//结构体实例
			worker := &MyWork{
				Name:      "workpool index " + strconv.Itoa(i),
				BirthYear: i,
				WP:        workerWorkPool,
			}
			err := workerWorkPool.PostWork("name_routine", worker)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err)
				time.Sleep(100 * time.Millisecond)
			}
			if shutdown == true {
				return
			}
		}
	}()

	fmt.Println("--------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	shutdown = true

	fmt.Println("Shutting Down\n")
	workerWorkPool.Shutdown("结束线程池 ！！！")
}
