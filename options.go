package dynamiclevelcache

import "time"

const (
	defaultStoreDurationgo = 5 * time.Minute
	defaultStoreLevelCount = 2
)

// StoreType : as name says
type StoreType int

// StoreType : as title says
const (
	StoreTypeNone StoreType = iota
	StoreTypeLevelAvg
	StoreTypeLevelWeight
)

// DurationType : as name says
type DurationType int

// DurationType : as name says
const (
	DurationTypeNone DurationType = iota
	DurationTypeFixed
	DurationTypeWeight
	DurationTypeDynamic
)

// DurationTypeOptions : as name says
type DurationTypeOptions struct {
	Type            DurationType
	Weights         []int
	WeightDurations []time.Duration
	FixedDuration   time.Duration
}

// StoreTypeOptions : as name says
type StoreTypeOptions struct {
	Type    StoreType
	Weights []int
}

// Options : init options
type Options struct {
	StoreOptions    StoreTypeOptions
	DurationOptions DurationTypeOptions
	CacheLevel      int
}

// NewDefaultOptions :
func NewDefaultOptions() Options {
	return Options{
		StoreOptions:    NewDefaultStoreOptions(),
		DurationOptions: NewDetaultDurationOptions(),
		CacheLevel:      defaultStoreLevelCount,
	}
}

// NewDefaultStoreOptions : as name says
func NewDefaultStoreOptions() StoreTypeOptions {
	return NewLevelAvgStoreOptions()
}

// NewLevelAvgStoreOptions : as name says
func NewLevelAvgStoreOptions() StoreTypeOptions {
	return StoreTypeOptions{
		Type: StoreTypeLevelAvg,
	}
}

// NewLevelWeightStoreOptions : as name says
func NewLevelWeightStoreOptions(weights []int) StoreTypeOptions {
	return StoreTypeOptions{
		Type:    StoreTypeLevelAvg,
		Weights: weights,
	}
}

// NewDetaultDurationOptions : as name says
func NewDetaultDurationOptions() DurationTypeOptions {
	return NewFixedDurationOptions(defaultStoreDurationgo)
}

// NewFixedDurationOptions : as name says
func NewFixedDurationOptions(d time.Duration) DurationTypeOptions {
	return DurationTypeOptions{
		Type:          DurationTypeFixed,
		FixedDuration: d,
	}
}

// NewWeightDurationOptions : as name says
func NewWeightDurationOptions(weights []int, durations []time.Duration) DurationTypeOptions {
	return DurationTypeOptions{
		Type:            DurationTypeWeight,
		Weights:         weights,
		WeightDurations: durations,
	}
}

// NewDynamicDurationOptions : as name says
func NewDynamicDurationOptions() DurationTypeOptions {
	return DurationTypeOptions{
		Type: DurationTypeDynamic,
	}
}
