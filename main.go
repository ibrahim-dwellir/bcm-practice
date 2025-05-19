package main

import (
	"fmt"
	"os"
	"time"

	"dwellir.com/bcm/database"
	"dwellir.com/bcm/job"
	"dwellir.com/bcm/logger"
	"dwellir.com/bcm/timer"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Job struct {
	Name     string `yaml:"name"`
	Endpoint string `yaml:"endpoint"`
	Interval *int   `yaml:"interval,omitempty"`
}

type Config struct {
	Interval int   `yaml:"interval"`
	Jobs     []Job `yaml:"jobs"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("not .env file found proceeding without it")
	}

	log := logger.GetLogger()

	// Load endpoint configuration
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	// Connect to the database
	conn, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Apply migrations
	database.ApplyMigration(conn)

	// Start a coroutine rest api request to each endpoint
	for _, jobProp := range cfg.Jobs {
		interval := cfg.Interval
		if jobProp.Interval != nil {
			interval = *jobProp.Interval
		}

		timer.NewInterval(func() {
			job.New(conn, jobProp.Name, jobProp.Endpoint)
		}, time.Duration(interval))
	}

	select {} // Keep the main goroutine alive
}
