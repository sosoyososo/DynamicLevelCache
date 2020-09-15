package dynamiclevelcache

import (
	"time"
)

// Fetcher : data fetcher
type Fetcher func() (interface{}, error)

// Saver : cache data saver
type Saver func(level int, key string, data interface{}, duration time.Duration) error

// Remover : cache data remover
type Remover func(key string, level int) error

// CacheDispatcher : as name says
type CacheDispatcher struct {
	opts Options
	s    Saver
	r    Remover
}

// DefaultDispatcher : as name says
func DefaultDispatcher(s Saver, r Remover) *CacheDispatcher {
	return &CacheDispatcher{
		opts: NewDefaultOptions(),
	}
}

// NewDispatcher : as name says
func NewDispatcher(opts Options, s Saver, r Remover) *CacheDispatcher {
	return &CacheDispatcher{
		opts: opts,
		s:    s,
		r:    r,
	}
}

// Get : get data from cache,
/**
 * if not found from cache , use f fetch data,
 * then result in suitable cache for next use
 * **/
func (d *CacheDispatcher) Get(key string, f Fetcher) (interface{}, error) {
	return nil, nil
}

// Remove : remove data from all cache
func (d *CacheDispatcher) Remove(key string) error {
	for i := 1; i <= d.opts.CacheLevel; i++ {
		err := d.r(key, i)
		if nil != err {
			return err
		}
	}
	return nil
}
