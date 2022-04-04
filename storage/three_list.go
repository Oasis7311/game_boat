package storage

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/cast"

	"github.com/oasis/game_boat/biz/dal/mysql/game_dal"
	utils2 "github.com/oasis/game_boat/utils"
)

var threeList = map[string][]uint{}
var lock = sync.RWMutex{}
var rawGameIds = make([]interface{}, 0)
var todayTime = time.Time{}
var nowTime = time.Time{}

func InitStorage() {
	go func() {
		for true {
			gameIds, err := game_dal.GetAllGamesIdList()
			if err != nil {
				panic(fmt.Sprintf("%+v", err))
			}

			for _, id := range gameIds {
				rawGameIds = append(rawGameIds, id)
			}

			nowTime = time.Now()

			todayTime = time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)

			appendList("new_game", todayTime.UnixNano())
			appendList("recent_update", todayTime.UnixNano()-17)
			appendList("hot_game", todayTime.Unix()-61)

			nextDayTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()+1, 0, 0, 5, 0, time.Local)
			waitTime := nextDayTime.Unix() - nowTime.Unix()
			time.Sleep(time.Duration(waitTime) * time.Second)
		}
	}()
}

func appendList(key string, seed int64) {
	RawGameIds := utils2.FakeShuffleNumSlice(rawGameIds, seed)
	lock.Lock()
	for _, id := range RawGameIds {
		threeList[key] = append(threeList[key], cast.ToUint(id))
	}
	lock.Unlock()

}

func GetThreeList() map[string][]uint {
	lock.RLock()
	defer lock.RUnlock()
	tmp := threeList
	return tmp
}
