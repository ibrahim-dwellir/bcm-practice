package job

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"dwellir.com/bcm/logger"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

var payload = map[string]interface{}{
	"id":      1,
	"jsonrpc": "2.0",
	"method":  "eth_blockNumber",
}

func New(conn driver.Conn, name string, endpoint string) {
	var log = logger.GetLogger()

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
		return
	}

	start := time.Now()
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	latency := time.Since(start)

	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
		return
	}

	blockHeightHex, ok := result["result"].(string)
	if !ok {
		log.Fatalf("Invalid response format: %v", result)
		return
	}

	blockHeight, err := strconv.ParseUint(blockHeightHex, 0, 64)
	if err != nil {
		log.Fatalf("Failed to parse block height: %v", err)
		return
	}

	log.Debugf("%s: blockHeight: %d\n", name, blockHeight)

	// Insert the block height into the database
	err = conn.Exec(context.Background(), `
			INSERT INTO block_heights (name, height, timestamp, latency_ms)
			VALUES (?, ?, now(), ?)
		`, name, blockHeight, latency.Milliseconds())

	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
		return
	}
}
