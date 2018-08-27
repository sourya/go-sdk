package stats

import "fmt"

// Assert that the mock collector implements Collector.
var (
	_ Collector = &MockCollector{}
)

// NewMockCollector returns a new mock collector.
func NewMockCollector() *MockCollector {
	return &MockCollector{
		Events: make(chan MockMetric),
	}
}

// MockCollector is a mocked collector for stats.
type MockCollector struct {
	namespace   string
	defaultTags []string

	Events chan MockMetric
}

// DefaultTags returns the default tags set.
func (mc MockCollector) DefaultTags() []string {
	return mc.defaultTags
}

// WithDefaultTag adds a default tag and returns a reference to the mock collector.
func (mc *MockCollector) WithDefaultTag(key, value string) *MockCollector {
	mc.defaultTags = append(mc.defaultTags, fmt.Sprintf("%s:%s", key, value))
	return mc
}

// Count adds a mock count event to the event stream.
func (mc *MockCollector) Count(name string, value int64, tags ...string) error {
	mc.Events <- MockMetric{Name: name, Count: value, Tags: append(mc.defaultTags, tags...)}
	return nil
}

// Increment adds a mock count event to the event stream with value (1).
func (mc *MockCollector) Increment(name string, tags ...string) error {
	mc.Events <- MockMetric{Name: name, Count: 1, Tags: append(mc.defaultTags, tags...)}
	return nil
}

// Gauge adds a mock count event to the event stream with value (1).
func (mc *MockCollector) Gauge(name string, value float64, tags ...string) error {
	mc.Events <- MockMetric{Name: name, Gauge: value, Tags: append(mc.defaultTags, tags...)}
	return nil
}

// Histogram adds a mock count event to the event stream with value (1).
func (mc *MockCollector) Histogram(name string, value float64, tags ...string) error {
	mc.Events <- MockMetric{Name: name, Histogram: value, Tags: append(mc.defaultTags, tags...)}
	return nil
}

// MockMetric is a mock metric.
type MockMetric struct {
	Name      string
	Count     int64
	Gauge     float64
	Histogram float64
	Tags      []string
}