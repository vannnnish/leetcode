package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个请求来了，把需要执行的业务逻辑封装成Task，放入木桶，等待worker取出执行
type Task struct {
	handler func() Result // worker从木桶中取出请求对象后要执行的业务逻辑函数
	resChan chan Result   // 等待worker执行并返回结果的channel
	taskID  int
}

// 封装业务逻辑的执行结果
type Result struct {
}

// 模拟业务逻辑的函数
func handler() Result {
	time.Sleep(300 * time.Millisecond)
	return Result{}
}

func NewTask(id int) Task {
	return Task{
		handler: handler,
		resChan: make(chan Result),
		taskID:  id,
	}
}

// 漏桶
type LeakyBucket struct {
	BucketSize int       // 木桶的大小
	NumWorker  int       // 同时从木桶中获取任务执行的worker数量
	bucket     chan Task // 存方任务的木桶
}

func NewLeakyBucket(bucketSize int, numWorker int) *LeakyBucket {
	return &LeakyBucket{
		BucketSize: bucketSize,
		NumWorker:  numWorker,
		bucket:     make(chan Task, bucketSize),
	}
}

func (b *LeakyBucket) validate(task Task) bool {
	// 如果木桶已经满了，返回false
	select {
	case b.bucket <- task:
	default:
		fmt.Printf("request[id=%d] is refused\n", task.taskID)
		return false
	}

	// 等待worker执行
	<-task.resChan
	fmt.Printf("request[id=%d] is run\n", task.taskID)
	return true
}

func (b *LeakyBucket) Start() {
	// 开启worker从木桶拉取任务执行
	go func() {
		// NumWorker 最大的并发数量
		// NumWorker 个goroutine 去执行BucketSize个任务
		for i := 0; i < b.NumWorker; i++ {
			go func() {
				for {
					// 取出一个任务
					task := <-b.bucket
					// 执行完成
					result := task.handler()
					// 将任务结果写入过resChan里面
					task.resChan <- result
				}
			}()
		}
	}()
}

func main() {
	// 创建一个漏桶
	bucket := NewLeakyBucket(10, 4)
	//
	bucket.Start()

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 创建一个任务
			task := NewTask(id)
			// 将这个任务添加到漏桶之中,并且执行
			bucket.validate(task)
		}(i)
	}
	wg.Wait()
}
