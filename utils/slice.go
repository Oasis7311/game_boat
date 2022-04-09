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
	count := 3000
	for count > 0 {
		randBase = append(randBase, rand.Int())
		count--
	}

	randBase = UniqueSlice(randBase)

	initRandBase := make([]interface{}, len(randBase))
	copy(initRandBase, randBase)

	randBaseMap := map[int]int{}
	newSlice = make([]interface{}, len(oldSlice))

	sort.Slice(randBase, func(i, j int) bool {
		return cast.ToInt(randBase[i]) < cast.ToInt(randBase[j])
	})

	for i, num := range randBase {
		randBaseMap[cast.ToInt(num)] = i
	}

	j := 0
	for _, num := range oldSlice {
		for j < len(initRandBase) && randBaseMap[cast.ToInt(initRandBase[j])] >= len(newSlice) {
			j++
		}
		newSlice[randBaseMap[cast.ToInt(initRandBase[j])]] = num
		j++
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
