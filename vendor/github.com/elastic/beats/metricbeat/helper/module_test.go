// +build !integration

package helper

import (
	"reflect"
	"testing"

	"github.com/elastic/beats/libbeat/common"

	"github.com/stretchr/testify/assert"
)

func TestGetMetricSetsList(t *testing.T) {

	metricSets := map[string]*MetricSet{}
	metricSets["test1"] = &MetricSet{}
	metricSets["test2"] = &MetricSet{}

	module := Module{
		metricSets: metricSets,
	}

	assert.Equal(t, "test1, test2", module.getMetricSetsList())

}

func TestNewModule(t *testing.T) {

	config, _ := common.NewConfigFrom(ModuleConfig{
		Module: "test",
	})

	module, err := NewModule(config, NewMockModuler)
	assert.NoError(t, err)
	assert.NotNil(t, module)

	err = module.moduler.Setup(module)
	assert.NoError(t, err)
}

// Check that the moduler inside each module is a different instance
func TestNewModulerDifferentInstance(t *testing.T) {

	config, _ := common.NewConfigFrom(ModuleConfig{
		Module: "test",
	})

	module1, err := NewModule(config, NewMockModuler)
	assert.NoError(t, err)
	module2, err := NewModule(config, NewMockModuler)
	assert.NoError(t, err)

	module1.moduler.Setup(module1)
	assert.False(t, reflect.DeepEqual(module1.moduler, module2.moduler))
	assert.True(t, reflect.DeepEqual(module1.moduler, module1.moduler))
}

// New creates new instance of Moduler
func NewMockModuler() Moduler {
	return &MockModuler{}
}

type MockModuler struct {
	counter int
}

func (m *MockModuler) Setup(mo *Module) error {
	m.counter += 1
	return nil
}
