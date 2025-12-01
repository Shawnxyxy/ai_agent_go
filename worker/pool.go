package worker

import (
	"ai_agent/agent"
	"ai_agent/model"
	"ai_agent/service"
	"fmt"
)

type Job struct {
	Name string
	Data interface{}
}

type WorkerPool struct {
	JobChan chan Job
}

func NewWorkerPool(workerCount int) *WorkerPool {
	pool := &WorkerPool{
		JobChan: make(chan Job, 100), // 队列的最大容量是 100 个任务
	}

	for i := 0; i < workerCount; i++ { // 并发执行 workerCount 个 任务
		go pool.worker(i)
	}
	return pool
}

func (p *WorkerPool) worker(id int) {
	for job := range p.JobChan {
		fmt.Printf("Worker %d is processing job: %s\n", id, job.Name)

		// 取出task数据
		task, ok := job.Data.(model.Task)
		if !ok {
			fmt.Println("Invalid task")
			continue
		}
		// 调用agent处理器
		result, err := agent.HandleTask(task)
		status := "success"
		if err != nil {
			status = "failed"
			result = err.Error()
		}
		// 更新数据库
		service.UpdateTaskStatus(task.ID, status, result)
	}
}

func (p *WorkerPool) AddJob(job Job) {
	p.JobChan <- job
}