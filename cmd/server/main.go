package main

import (
	"flag"
	"log"
	"snake-game/internal/cache"
	"snake-game/internal/database"
	"snake-game/internal/game"
	"snake-game/internal/kafka"
	"snake-game/internal/network"
	"snake-game/internal/pool"
	"snake-game/pkg/config"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewMySQLDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	redisClient, err := cache.NewRedisClient(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	kafkaProducer, err := kafka.NewProducer(cfg.KafkaURL)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	kafkaConsumer, err := kafka.NewConsumer(cfg.KafkaURL, "game-events")
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer kafkaConsumer.Close()

	workerPool := pool.NewWorkerPool(cfg.WorkerPoolSize)
	defer workerPool.Close()

	gameInstance := game.NewGame()

	// 创建服务器实例，传入必要的参数
	server, err := network.NewServer(gameInstance, cfg.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// 启动服务器
	log.Printf("Starting server on %s", cfg.ServerAddress)
	server.Start() // 修正：去掉对返回值的赋值
}
