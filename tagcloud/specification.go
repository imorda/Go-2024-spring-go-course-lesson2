package tagcloud

import "sort"

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	storage map[string]int
}

// TagStat represents statistics regarding single tag
type TagStat struct {
	Tag             string
	OccurrenceCount int
}

// New should create a valid TagCloud instance
func New() *TagCloud {
	return &TagCloud{
		storage: make(map[string]int),
	}
}

// AddTag should add a tag to the cloud if it wasn't present and increase tag occurrence count
// thread-safety is not needed
func (tc *TagCloud) AddTag(tagName string) {
	tc.storage[tagName]++
}

// TopN should return top N most frequent tags ordered in descending order by occurrence count
// if there are multiple tags with the same occurrence count then the order is defined by implementation
// if n is greater that TagCloud size then all elements should be returned
// thread-safety is not needed
// there are no restrictions on time complexity
func (tc *TagCloud) TopN(n int) []TagStat {
	result := make([]TagStat, 0, len(tc.storage))

	for k, v := range tc.storage {
		result = append(result, TagStat{
			Tag:             k,
			OccurrenceCount: v,
		})
	}

	sort.Slice(result, func(lhs, rhs int) bool {
		return result[lhs].OccurrenceCount > result[rhs].OccurrenceCount
	})

	return result[:min(len(result), n)]
}
