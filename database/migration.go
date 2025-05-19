package database

import (
	"context"

	"dwellir.com/bcm/logger"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

var log = logger.GetLogger()

func ApplyMigration(conn driver.Conn) {
	log.Info("Applying migration...")
	createBlockHeightTable(conn)
	log.Info("Migration applied successfully.")
}

func createBlockHeightTable(conn driver.Conn) {
	err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS block_heights (
			name String,
			height UInt64,
			timestamp DateTime64(3, 'UTC')
		) ENGINE = MergeTree()
		PARTITION BY toYYYYMM(timestamp)
		ORDER BY timestamp
	`)

	if err != nil {
		log.Fatalf("Failed to create table block_heights: %v", err)
	}
}
