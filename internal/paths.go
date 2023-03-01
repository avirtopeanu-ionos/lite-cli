package paths

import (
	"sort"
	"strings"
)

func SortByPartCount(paths []string) []string {
	sort.Slice(paths, func(i, j int) bool {
		parts := func(idx int) int {
			return strings.Count(paths[idx], "/") - strings.Count(paths[idx], "{")
		}
		if parts(i) == parts(j) {
			return strings.Count(paths[i], "{") < strings.Count(paths[j], "{")
		}
		return parts(i) < parts(j)
	})
	return paths
}
