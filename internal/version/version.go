package version

import (
	"encoding/json"
	"fmt"
)

// These variables should be set during app build
var (
	buildDate = "UNKNOWN"
	gitHash   = "UNKNOWN"
)

type AppVersion struct {
	BuildDate string `json:"build_date"`
	GitHash   string `json:"gitHash"`
}

func GetCurrentVersion() (string, error) {
	ver := AppVersion{
		BuildDate: buildDate,
		GitHash:   gitHash,
	}
	result, err := json.Marshal(ver)
	if err != nil {
		return "", fmt.Errorf("marshal version to json: %w", err)
	}
	return string(result), nil
}
