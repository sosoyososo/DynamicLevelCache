package dynamiclevelcache

import (
	"time"
)

// StepType :  data oprtation step
type StepType int

// StepType :  data oprtation step
const (
	StepTypeNone StepType = iota
	StepTypeGetFromCache
	StepTypeSaveToCache
	StepTypeRemoveFromCache
	StepTypeFetchFromOrignalSource
	StepTypeInDispatching
)

// Fetcher : data fetcher
type Fetcher func() (interface{}, error)

// Saver : cache data saver
type Saver func(level int, key string, du time.Duration, data interface{}) error

// Getter : cache data getter
type Getter func(level int, key string) (interface{}, error)

// Remover : cache data remover
type Remover func(key string, level int) error

// FailureCallBack : data handle failing callback
type FailureCallBack func(step StepType, err error)

// CacheDispatcher : as name says
type CacheDispatcher struct {
	opts Options
	s    Saver
	r    Remover
	g    Getter
	f    FailureCallBack
}

// DefaultDispatcher : panic if any parameter is nil
func DefaultDispatcher(s Saver, r Remover, g Getter, f FailureCallBack) *CacheDispatcher {
	if nil == s || nil == r || nil == g {
		panic("wrong input parameter")
	}

	return &CacheDispatcher{
		opts: NewDefaultOptions(),
		s:    s,
		g:    g,
		r:    r,
		f:    f,
	}
}

// NewDispatcher : panic if any parameter is nil or wrong options
func NewDispatcher(opts Options, s Saver, r Remover, g Getter, f FailureCallBack) *CacheDispatcher {
	if nil == s || nil == r || nil == g {
		panic("wrong input parameter")
	}
	if StoreTypeNone == opts.StoreOptions.Type || DurationTypeNone == opts.DurationOptions.Type {
		panic("wrong input options")
	}

	return &CacheDispatcher{
		opts: opts,
		s:    s,
		g:    g,
		r:    r,
	}
}

// Get : get data from cache,
/**
 * if not found from cache , use f fetch data,
 * then result in suitable cache for next use
 * **/
func (d *CacheDispatcher) Get(key string, f Fetcher) (interface{}, error) {
	// get data from cache
	data, err := d.get(key)
	if nil != err || nil != data {
		return data, err
	}

	// fetch data from orignal source
	data, err = d.fetch(key, f)
	defer func() {
		if nil == err {
			// save data to cache source
			err = d.save(key, data)
			if nil != err {
				stepFail(d, StepTypeSaveToCache, err)
			}
		}
	}()
	return data, err
}

// Remove : remove data from  cache
func (d *CacheDispatcher) Remove(key string) error {
	return d.remove(key)
}
