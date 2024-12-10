package utils

import (
	"sort"
	"strings"
	"time"

	"github.com/Yahya-Elamri/signeitfaster/module"
)

func parseTimestamp(timestamp []uint8) (time.Time, error) {
	// Convert []uint8 to string
	timestampStr := string(timestamp)

	// Parse the string into a time.Time object
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timestampStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func SortByTimestamp(items []module.MixedListing) ([]module.MixedListing, error) {
	// Sort the slice in-place by converting the timestamp and comparing
	sort.Slice(items, func(i, j int) bool {
		timeI, errI := parseTimestamp(items[i].CreatedAt)
		timeJ, errJ := parseTimestamp(items[j].CreatedAt)

		// Handle potential parsing errors
		if errI != nil || errJ != nil {
			return false
		}

		return timeI.After(timeJ)
	})

	// Return the sorted slice
	return items, nil
}

func Transform(jsonString string) []string {
	cleanedString := strings.Trim(jsonString, "[]")
	cleanedString = strings.ReplaceAll(cleanedString, `"`, "")
	paths := strings.Split(cleanedString, ",")
	return paths
}
