package medium

type tuple struct {
	value     string
	timestamp int
}

type TimeMap struct {
	hashTable map[string][]tuple
}

func Constructor() TimeMap {
	return TimeMap{
		hashTable: make(map[string][]tuple),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.hashTable[key] = append(this.hashTable[key], tuple{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.hashTable[key]; !ok {
		return ""
	}

	l, r := 0, len(this.hashTable[key])-1
	for l < r {
		m := (l + r + 1) / 2
		if this.hashTable[key][m].timestamp <= timestamp {
			l = m
		} else {
			r = m - 1
		}
	}

	if this.hashTable[key][l].timestamp > timestamp {
		return ""
	}

	return this.hashTable[key][l].value
}
