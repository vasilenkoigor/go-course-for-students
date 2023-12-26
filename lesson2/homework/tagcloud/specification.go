package tagcloud

import (
	"sort"
)

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	// Private
	storage []TagStat
}

// TagStat represents statistics regarding single tag
type TagStat struct {
	// Public
	Tag             string
	OccurrenceCount int
}

// New should create a valid TagCloud instance
func New() TagCloud {
	return TagCloud{storage: make([]TagStat, 0)}
}

// AddTag should add a tag to the cloud if it wasn't present and increase tag occurrence count
// thread-safety is not needed
func (cloud *TagCloud) AddTag(tag string) {
	var tagStat TagStat
	var tagStatIndex int
	var hasValue bool

	for index, stat := range cloud.storage {
		if stat.Tag == tag {
			tagStat = stat
			tagStatIndex = index
			hasValue = true
			break
		}
	}

	if hasValue {
		tagStat = TagStat{Tag: tag, OccurrenceCount: tagStat.OccurrenceCount + 1}
		cloud.storage[tagStatIndex] = tagStat
	} else {
		tagStat = TagStat{Tag: tag, OccurrenceCount: 1}
		cloud.storage = append(cloud.storage, tagStat)
	}

	sort.Slice(cloud.storage, func(i, j int) bool {
		return cloud.storage[i].OccurrenceCount > cloud.storage[j].OccurrenceCount
	})
}

// TopN should return top N most frequent tags ordered in descending order by occurrence count
// if there are multiple tags with the same occurrence count then the order is defined by implementation
// if n is greater that TagCloud size then all elements should be returned
// thread-safety is not needed
// there are no restrictions on time complexity
func (cloud *TagCloud) TopN(n int) []TagStat {
	if n > len(cloud.storage) {
		return cloud.storage
	} else {
		return cloud.storage[0:n]
	}
}
