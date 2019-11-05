package dice

import (
	"fmt"
	"strings"
)

// This file will include common functions shared amongst the package files

// Returns a key value map with the valies retrieved from a (grouped) regex
func getMapRegexGroups(keys []string, values []string) map[string]string {
	var groups = map[string]string{}
	for i, value := range values {
		if keys[i] != "" {
			groups[keys[i]] = value
		} else {
			groups["command"] = value
		}
	}
	return groups
}

// Returns the addition of a slice's items
func getSliceItemsTotal(slice []int) (total int) {
	for _, value := range slice {
		total += value
	}
	return
}

// Returns a string with slice values separated by commas
func getStringValues(slice []int) string {
	var str strings.Builder
	for _, value := range slice {
		fmt.Fprintf(&str, "%d,", value)
	}
	result := strings.TrimSuffix(str.String(), ",")
	return fmt.Sprintf("%s", result)
}
