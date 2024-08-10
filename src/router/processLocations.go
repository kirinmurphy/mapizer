package router

import (
	"context"
	"encoding/json"
	"fmt"
	"gomap/src/locationManager"

	"github.com/redis/go-redis/v9"
)

func processLocations(sheetId string, redisClient *redis.Client, ctx context.Context) error {
	spreadsheetUrl := fmt.Sprintf(baseSpreadsheetUrl, sheetId)

	parsedLocations, err := locationManager.LoadLocations(spreadsheetUrl)
	if err != nil {
		return fmt.Errorf("failed to load locations: %w", err)
	}

	locationsJson, err := json.Marshal(parsedLocations)
	if err != nil {
		return fmt.Errorf("failed to marshal locations: %w", err)
	}

	err = redisClient.Set(ctx, sheetId, locationsJson, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to cache locations: %w", err)
	}

	return nil
}
