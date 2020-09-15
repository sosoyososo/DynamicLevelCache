package dynamiclevelcache

import (
	"time"
)

/***
 *
 * cache saver use this service to determine which level and how long to store data
 *
 * */
func (d *CacheDispatcher) dataStoreStoreLevelAndDuration(key string) (int, time.Duration) {
	return 0, 0
}
