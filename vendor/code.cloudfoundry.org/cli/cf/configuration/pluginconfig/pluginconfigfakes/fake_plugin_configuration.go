// This file was generated by counterfeiter
package pluginconfigfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/configuration/pluginconfig"
)

type FakePluginConfiguration struct {
	PluginsStub        func() map[string]pluginconfig.PluginMetadata
	pluginsMutex       sync.RWMutex
	pluginsArgsForCall []struct{}
	pluginsReturns     struct {
		result1 map[string]pluginconfig.PluginMetadata
	}
	SetPluginStub        func(string, pluginconfig.PluginMetadata)
	setPluginMutex       sync.RWMutex
	setPluginArgsForCall []struct {
		arg1 string
		arg2 pluginconfig.PluginMetadata
	}
	GetPluginPathStub        func() string
	getPluginPathMutex       sync.RWMutex
	getPluginPathArgsForCall []struct{}
	getPluginPathReturns     struct {
		result1 string
	}
	RemovePluginStub        func(string)
	removePluginMutex       sync.RWMutex
	removePluginArgsForCall []struct {
		arg1 string
	}
	ListCommandsStub        func() []string
	listCommandsMutex       sync.RWMutex
	listCommandsArgsForCall []struct{}
	listCommandsReturns     struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePluginConfiguration) Plugins() map[string]pluginconfig.PluginMetadata {
	fake.pluginsMutex.Lock()
	fake.pluginsArgsForCall = append(fake.pluginsArgsForCall, struct{}{})
	fake.recordInvocation("Plugins", []interface{}{})
	fake.pluginsMutex.Unlock()
	if fake.PluginsStub != nil {
		return fake.PluginsStub()
	} else {
		return fake.pluginsReturns.result1
	}
}

func (fake *FakePluginConfiguration) PluginsCallCount() int {
	fake.pluginsMutex.RLock()
	defer fake.pluginsMutex.RUnlock()
	return len(fake.pluginsArgsForCall)
}

func (fake *FakePluginConfiguration) PluginsReturns(result1 map[string]pluginconfig.PluginMetadata) {
	fake.PluginsStub = nil
	fake.pluginsReturns = struct {
		result1 map[string]pluginconfig.PluginMetadata
	}{result1}
}

func (fake *FakePluginConfiguration) SetPlugin(arg1 string, arg2 pluginconfig.PluginMetadata) {
	fake.setPluginMutex.Lock()
	fake.setPluginArgsForCall = append(fake.setPluginArgsForCall, struct {
		arg1 string
		arg2 pluginconfig.PluginMetadata
	}{arg1, arg2})
	fake.recordInvocation("SetPlugin", []interface{}{arg1, arg2})
	fake.setPluginMutex.Unlock()
	if fake.SetPluginStub != nil {
		fake.SetPluginStub(arg1, arg2)
	}
}

func (fake *FakePluginConfiguration) SetPluginCallCount() int {
	fake.setPluginMutex.RLock()
	defer fake.setPluginMutex.RUnlock()
	return len(fake.setPluginArgsForCall)
}

func (fake *FakePluginConfiguration) SetPluginArgsForCall(i int) (string, pluginconfig.PluginMetadata) {
	fake.setPluginMutex.RLock()
	defer fake.setPluginMutex.RUnlock()
	return fake.setPluginArgsForCall[i].arg1, fake.setPluginArgsForCall[i].arg2
}

func (fake *FakePluginConfiguration) GetPluginPath() string {
	fake.getPluginPathMutex.Lock()
	fake.getPluginPathArgsForCall = append(fake.getPluginPathArgsForCall, struct{}{})
	fake.recordInvocation("GetPluginPath", []interface{}{})
	fake.getPluginPathMutex.Unlock()
	if fake.GetPluginPathStub != nil {
		return fake.GetPluginPathStub()
	} else {
		return fake.getPluginPathReturns.result1
	}
}

func (fake *FakePluginConfiguration) GetPluginPathCallCount() int {
	fake.getPluginPathMutex.RLock()
	defer fake.getPluginPathMutex.RUnlock()
	return len(fake.getPluginPathArgsForCall)
}

func (fake *FakePluginConfiguration) GetPluginPathReturns(result1 string) {
	fake.GetPluginPathStub = nil
	fake.getPluginPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginConfiguration) RemovePlugin(arg1 string) {
	fake.removePluginMutex.Lock()
	fake.removePluginArgsForCall = append(fake.removePluginArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RemovePlugin", []interface{}{arg1})
	fake.removePluginMutex.Unlock()
	if fake.RemovePluginStub != nil {
		fake.RemovePluginStub(arg1)
	}
}

func (fake *FakePluginConfiguration) RemovePluginCallCount() int {
	fake.removePluginMutex.RLock()
	defer fake.removePluginMutex.RUnlock()
	return len(fake.removePluginArgsForCall)
}

func (fake *FakePluginConfiguration) RemovePluginArgsForCall(i int) string {
	fake.removePluginMutex.RLock()
	defer fake.removePluginMutex.RUnlock()
	return fake.removePluginArgsForCall[i].arg1
}

func (fake *FakePluginConfiguration) ListCommands() []string {
	fake.listCommandsMutex.Lock()
	fake.listCommandsArgsForCall = append(fake.listCommandsArgsForCall, struct{}{})
	fake.recordInvocation("ListCommands", []interface{}{})
	fake.listCommandsMutex.Unlock()
	if fake.ListCommandsStub != nil {
		return fake.ListCommandsStub()
	} else {
		return fake.listCommandsReturns.result1
	}
}

func (fake *FakePluginConfiguration) ListCommandsCallCount() int {
	fake.listCommandsMutex.RLock()
	defer fake.listCommandsMutex.RUnlock()
	return len(fake.listCommandsArgsForCall)
}

func (fake *FakePluginConfiguration) ListCommandsReturns(result1 []string) {
	fake.ListCommandsStub = nil
	fake.listCommandsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakePluginConfiguration) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pluginsMutex.RLock()
	defer fake.pluginsMutex.RUnlock()
	fake.setPluginMutex.RLock()
	defer fake.setPluginMutex.RUnlock()
	fake.getPluginPathMutex.RLock()
	defer fake.getPluginPathMutex.RUnlock()
	fake.removePluginMutex.RLock()
	defer fake.removePluginMutex.RUnlock()
	fake.listCommandsMutex.RLock()
	defer fake.listCommandsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePluginConfiguration) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ pluginconfig.PluginConfiguration = new(FakePluginConfiguration)
