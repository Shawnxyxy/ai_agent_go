package scheduler

import (
	"ai_agent/service"
	"ai_agent/worker"
	"time"
)

func StartScheduler(pool *worker.WorkerPool) {
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		tasks, _ := service.GetPendingTasks()

		for _, task := range tasks {
			pool.AddJob(worker.Job{
				Name: task.Type,
				Data: task,
			})

			service.UpdateTaskStatus(task.ID, "running", "")
		}
	}
}