package utils

import (
	"math/rand"
	"sort"

	"github.com/spf13/cast"
)

// FakeShuffleNumSlice 伪随机洗牌
func FakeShuffleNumSlice(oldSlice []interface{}, seed int64) (newSlice []interface{}) {
	rand.Seed(seed)

	randBase := make([]interface{}, 0)
	count := 500
	for count > 0 {
		randBase = append(randBase, rand.Uint64())
		count--
	}

	randBase = UniqueSlice(randBase)
	initRandBase := randBase

	randBaseMap := map[uint64]int{}
	newSlice = make([]interface{}, len(oldSlice))

	sort.Slice(randBase, func(i, j int) bool {
		return cast.ToUint64(randBase[i]) < cast.ToUint64(randBase[j])
	})

	for i, num := range randBase {
		randBaseMap[cast.ToUint64(num)] = i
	}

	for i, num := range oldSlice {
		newSlice[randBaseMap[cast.ToUint64(initRandBase[i])]] = num
	}

	return newSlice
}

func UniqueSlice(arr []interface{}) []interface{} {
	set := make(map[interface{}]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}
