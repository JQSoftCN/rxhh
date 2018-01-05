package reports

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/muesli/cache2go"
	"time"
	"strings"
)

const (
	rconfs_cache_name = "rconfsCache"
	rconf_life_span   = time.Hour
)

var rconfsCache *cache2go.CacheTable

func initRConfCache() {
	rconfsCache = cache2go.Cache(rconfs_cache_name)
	rconfsCache.SetDataLoader(func(key interface{}, args ...interface{}) *cache2go.CacheItem {
		rname, _ := key.(string)
		calcCells := readConf(rname)
		rname = strings.ToLower(rname)
		item := cache2go.NewCacheItem(rname, rconf_life_span, calcCells)
		return item
	})
}

func GetRConf(rname string) *[]CalcCell {
	rname = strings.ToLower(rname)
	item, _ := rconfsCache.Value(rname)
	calcCells, _ := (item.Data()).(*[]CalcCell)
	return calcCells
}

func UpdateRConf(rname string) {
	rname = strings.ToLower(rname)
	rconfsCache.Delete(rname)
}

func readConf(rname string) *[]CalcCell {
	dir := RtDir(rname)
	bs, err := ioutil.ReadFile(dir + "/conf.json")
	if err != nil {
		log.Println(err)
	}

	var fmap []CalcCell
	err = json.Unmarshal(bs, &fmap)
	if err != nil {
		log.Println(err)
	}
	return &fmap
}
