/**
 * @Author: chentong
 * @Date: 2025/01/06 12:30
 */

package tools

import (
	"sync"
)

// WorkerPool 工作池结构
type WorkerPool struct {
	workerNum int            // 工作池中的 Worker 数量
	jobsChan  chan func()    // 任务通道，接收需要执行的任务
	wg        sync.WaitGroup // 用于等待所有 Worker 完成
}

// NewWorkerPool 创建一个新的工作池
func NewWorkerPool(workerNum, jobQueueSize int) *WorkerPool {
	return &WorkerPool{
		workerNum: workerNum,
		jobsChan:  make(chan func(), jobQueueSize),
	}
}

// Start 启动工作池
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerNum; i++ {
		wp.wg.Add(1)
		go func(workerID int) {
			defer wp.wg.Done()
			for job := range wp.jobsChan {
				job() // 执行任务
			}
		}(i + 1)
	}
}

// AddJob 向工作池中添加任务
func (wp *WorkerPool) AddJob(job func()) {
	wp.jobsChan <- job
}

// Wait 等待所有任务完成并关闭工作池
func (wp *WorkerPool) Wait() {
	close(wp.jobsChan) // 关闭任务通道，表示没有更多任务
	wp.wg.Wait()       // 等待所有 Worker 完成
}
