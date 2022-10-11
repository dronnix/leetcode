package main

import (
	"fmt"
	"sort"
)

type TimeMap struct {
	m map[string][]timedValue
}

type timedValue struct {
	val string
	ts  int
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]timedValue)}
}

func (tm *TimeMap) Set(key, value string, timestamp int) {
	nv := timedValue{val: value, ts: timestamp}
	tm.m[key] = append(tm.m[key], nv)
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	tvs := tm.m[key]
	l := len(tvs)
	if l == 0 {
		return ""
	}

	if tvs[l-1].ts < timestamp {
		return tvs[l-1].val
	}

	idx := sort.Search(len(tvs), func(i int) bool {
		return tvs[i].ts > timestamp
	})

	if idx == 0 {
		return ""
	}

	return tvs[idx-1].val
}

func main() {
	tm := Constructor()
	tm.Set("foo", "bar", 1)
	fmt.Println(tm.Get("foo", 1))
	fmt.Println(tm.Get("foo", 3))
	tm.Set("foo", "bar2", 4)
	fmt.Println(tm.Get("foo", 3))
	fmt.Println(tm.Get("foo", 5))
}
