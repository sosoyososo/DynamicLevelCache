package dynamiclevelcache

import "time"

func (d *CacheDispatcher) save(key string, data interface{}) error {
	l, du := d.dataStoreStoreLevelAndDuration(key)
	return d.s(l, key, du, data)
}

func (d *CacheDispatcher) get(key string) (interface{}, error) {
	for i := 1; i < d.opts.CacheLevel; i++ {
		data, err := d.g(i, key)
		if nil != err {
			return nil, err
		}
		if nil != data {
			return data, nil
		}
	}
	return nil, nil
}

func (d *CacheDispatcher) fetch(key string, f Fetcher) (interface{}, error) {
	start := time.Now()
	defer func() {
		offset := time.Now().Sub(start)
		d.dataFetched(key, start, offset)
	}()
	return f()
}

func (d *CacheDispatcher) remove(key string) error {
	for i := 1; i <= d.opts.CacheLevel; i++ {
		err := d.r(key, i)
		if nil != err {
			stepFail(d, StepTypeRemoveFromCache, err)
			return err
		}
	}
	return nil
}
