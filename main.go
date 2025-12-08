package main

import (
	"ai_agent/worker"
	"ai_agent/scheduler"
	"ai_agent/config"
	"ai_agent/database"
	"ai_agent/router"
	"ai_agent/model"
	"ai_agent/agent"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //  _ 表示 只导入包，但不直接使用其中的标识符
	"github.com/redis/go-redis/v9"
)

// 用于统一验证 App、MySQL、Redis 配置
func ValidateConfig() {
	fmt.Println("===== Config Validation Start =====")
	// App配置
	fmt.Println("App Config:")
	fmt.Println("Port:%d\n", config.Cfg.App.Port)
	fmt.Println("Other fields: %+v\n", config.Cfg.App)
	// MySQL配置
	fmt.Println("MySQL Config:")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Cfg.MySQL.User,
		config.Cfg.MySQL.Password,
		config.Cfg.MySQL.Host,
		config.Cfg.MySQL.Port,
		config.Cfg.MySQL.Database,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("MySQL open connection failed:", err)
	} else if err := db.Ping(); err != nil {
		fmt.Println("MySQL ping failed:", err)
	} else {
		fmt.Println("MySQL connection successful!")
	}
	if db != nil {
		db.Close()
	}

	// Redis配置
	fmt.Println("Redis Config:")
	rdb := redis.NewClient(&redis.Options{
		Addr:		fmt.Sprintf("%s:%d", config.Cfg.Redis.Host, config.Cfg.Redis.Port),
		Password:	config.Cfg.Redis.Password,
		DB:			config.Cfg.Redis.DB,
	})
	ctx := context.Background()
	testKey := "validate_test"
	err = rdb.Set(ctx, testKey, "ok", 0).Err()
	if err != nil {
		fmt.Println("Redis set failed:", err)
	} else if val, err := rdb.Get(ctx, testKey).Result(); err != nil {
		fmt.Println("Redis get failed:", err)
	} else {
		fmt.Println("Redis test value:", val)
	}
	fmt.Println("===== Config Validation End =====")
}

func main() {
	// 初始化配置
	config.InitConfig()
	// 初始化MySQL
	database.InitMySQL()
	// 初始化Redis
	database.InitRedis()

	// ----------- 临时测试 agent -----------
	fmt.Println("========== Start Agent Test ==========")
	testTasks := []model.Task{
		{Type: "echo", Payload: "Hello agent!"},
		{Type: "summarize", Payload: "Artificial Intelligence is transforming work and life."},
		{Type: "chat", Payload: "Hi, how are you?"},
		{Type: "recycle_memory"},
	}

	for _, t := range testTasks {
		fmt.Println("====================================")
		fmt.Println("Processing task:", t.Type)
		result, err := agent.HandleTask(t)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}
	fmt.Println("========== End Agent Test ==========")

	// 启动 worker pool
	pool := worker.NewWorkerPool(3) // 并发执行3个任务
	pool.AddJob(worker.Job{Name: "Warmup Embedding Model"})
	pool.AddJob(worker.Job{Name: "Initialize Memory"})
	// 启动定时调度器
	go scheduler.StartScheduler(pool)
	// 启动路由
	r := router.SetupRouter()
	// 从配置读取端口
	port := config.Cfg.App.Port
	r.Run(fmt.Sprintf(":%d", port))
}
