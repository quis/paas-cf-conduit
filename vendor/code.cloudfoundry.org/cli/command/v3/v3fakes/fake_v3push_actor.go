// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeV3PushActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	CreatePackageByApplicationNameAndSpaceStub        func(appName string, spaceGUID string, bitsPath string, dockerImageCredentials v3action.DockerImageCredentials) (v3action.Package, v3action.Warnings, error)
	createPackageByApplicationNameAndSpaceMutex       sync.RWMutex
	createPackageByApplicationNameAndSpaceArgsForCall []struct {
		appName                string
		spaceGUID              string
		bitsPath               string
		dockerImageCredentials v3action.DockerImageCredentials
	}
	createPackageByApplicationNameAndSpaceReturns struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}
	createPackageByApplicationNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}
	CreateApplicationInSpaceStub        func(app v3action.Application, spaceGUID string) (v3action.Application, v3action.Warnings, error)
	createApplicationInSpaceMutex       sync.RWMutex
	createApplicationInSpaceArgsForCall []struct {
		app       v3action.Application
		spaceGUID string
	}
	createApplicationInSpaceReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	createApplicationInSpaceReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	GetApplicationByNameAndSpaceStub        func(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	GetApplicationSummaryByNameAndSpaceStub        func(appName string, spaceGUID string) (v3action.ApplicationSummary, v3action.Warnings, error)
	getApplicationSummaryByNameAndSpaceMutex       sync.RWMutex
	getApplicationSummaryByNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
	}
	getApplicationSummaryByNameAndSpaceReturns struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}
	getApplicationSummaryByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}
	GetStreamingLogsForApplicationByNameAndSpaceStub        func(appName string, spaceGUID string, client v3action.NOAAClient) (<-chan *v3action.LogMessage, <-chan error, v3action.Warnings, error)
	getStreamingLogsForApplicationByNameAndSpaceMutex       sync.RWMutex
	getStreamingLogsForApplicationByNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
		client    v3action.NOAAClient
	}
	getStreamingLogsForApplicationByNameAndSpaceReturns struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}
	getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}
	PollStartStub        func(appGUID string, warnings chan<- v3action.Warnings) error
	pollStartMutex       sync.RWMutex
	pollStartArgsForCall []struct {
		appGUID  string
		warnings chan<- v3action.Warnings
	}
	pollStartReturns struct {
		result1 error
	}
	pollStartReturnsOnCall map[int]struct {
		result1 error
	}
	SetApplicationDropletStub        func(appName string, spaceGUID string, dropletGUID string) (v3action.Warnings, error)
	setApplicationDropletMutex       sync.RWMutex
	setApplicationDropletArgsForCall []struct {
		appName     string
		spaceGUID   string
		dropletGUID string
	}
	setApplicationDropletReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	setApplicationDropletReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	StagePackageStub        func(packageGUID string, appName string) (<-chan v3action.Droplet, <-chan v3action.Warnings, <-chan error)
	stagePackageMutex       sync.RWMutex
	stagePackageArgsForCall []struct {
		packageGUID string
		appName     string
	}
	stagePackageReturns struct {
		result1 <-chan v3action.Droplet
		result2 <-chan v3action.Warnings
		result3 <-chan error
	}
	stagePackageReturnsOnCall map[int]struct {
		result1 <-chan v3action.Droplet
		result2 <-chan v3action.Warnings
		result3 <-chan error
	}
	StartApplicationStub        func(appGUID string) (v3action.Application, v3action.Warnings, error)
	startApplicationMutex       sync.RWMutex
	startApplicationArgsForCall []struct {
		appGUID string
	}
	startApplicationReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	startApplicationReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	StopApplicationStub        func(appGUID string) (v3action.Warnings, error)
	stopApplicationMutex       sync.RWMutex
	stopApplicationArgsForCall []struct {
		appGUID string
	}
	stopApplicationReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	stopApplicationReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	UpdateApplicationStub        func(app v3action.Application) (v3action.Application, v3action.Warnings, error)
	updateApplicationMutex       sync.RWMutex
	updateApplicationArgsForCall []struct {
		app v3action.Application
	}
	updateApplicationReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	updateApplicationReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3PushActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeV3PushActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3PushActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3PushActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3PushActor) CreatePackageByApplicationNameAndSpace(appName string, spaceGUID string, bitsPath string, dockerImageCredentials v3action.DockerImageCredentials) (v3action.Package, v3action.Warnings, error) {
	fake.createPackageByApplicationNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.createPackageByApplicationNameAndSpaceReturnsOnCall[len(fake.createPackageByApplicationNameAndSpaceArgsForCall)]
	fake.createPackageByApplicationNameAndSpaceArgsForCall = append(fake.createPackageByApplicationNameAndSpaceArgsForCall, struct {
		appName                string
		spaceGUID              string
		bitsPath               string
		dockerImageCredentials v3action.DockerImageCredentials
	}{appName, spaceGUID, bitsPath, dockerImageCredentials})
	fake.recordInvocation("CreatePackageByApplicationNameAndSpace", []interface{}{appName, spaceGUID, bitsPath, dockerImageCredentials})
	fake.createPackageByApplicationNameAndSpaceMutex.Unlock()
	if fake.CreatePackageByApplicationNameAndSpaceStub != nil {
		return fake.CreatePackageByApplicationNameAndSpaceStub(appName, spaceGUID, bitsPath, dockerImageCredentials)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createPackageByApplicationNameAndSpaceReturns.result1, fake.createPackageByApplicationNameAndSpaceReturns.result2, fake.createPackageByApplicationNameAndSpaceReturns.result3
}

func (fake *FakeV3PushActor) CreatePackageByApplicationNameAndSpaceCallCount() int {
	fake.createPackageByApplicationNameAndSpaceMutex.RLock()
	defer fake.createPackageByApplicationNameAndSpaceMutex.RUnlock()
	return len(fake.createPackageByApplicationNameAndSpaceArgsForCall)
}

func (fake *FakeV3PushActor) CreatePackageByApplicationNameAndSpaceArgsForCall(i int) (string, string, string, v3action.DockerImageCredentials) {
	fake.createPackageByApplicationNameAndSpaceMutex.RLock()
	defer fake.createPackageByApplicationNameAndSpaceMutex.RUnlock()
	return fake.createPackageByApplicationNameAndSpaceArgsForCall[i].appName, fake.createPackageByApplicationNameAndSpaceArgsForCall[i].spaceGUID, fake.createPackageByApplicationNameAndSpaceArgsForCall[i].bitsPath, fake.createPackageByApplicationNameAndSpaceArgsForCall[i].dockerImageCredentials
}

func (fake *FakeV3PushActor) CreatePackageByApplicationNameAndSpaceReturns(result1 v3action.Package, result2 v3action.Warnings, result3 error) {
	fake.CreatePackageByApplicationNameAndSpaceStub = nil
	fake.createPackageByApplicationNameAndSpaceReturns = struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) CreatePackageByApplicationNameAndSpaceReturnsOnCall(i int, result1 v3action.Package, result2 v3action.Warnings, result3 error) {
	fake.CreatePackageByApplicationNameAndSpaceStub = nil
	if fake.createPackageByApplicationNameAndSpaceReturnsOnCall == nil {
		fake.createPackageByApplicationNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Package
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.createPackageByApplicationNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Package
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) CreateApplicationInSpace(app v3action.Application, spaceGUID string) (v3action.Application, v3action.Warnings, error) {
	fake.createApplicationInSpaceMutex.Lock()
	ret, specificReturn := fake.createApplicationInSpaceReturnsOnCall[len(fake.createApplicationInSpaceArgsForCall)]
	fake.createApplicationInSpaceArgsForCall = append(fake.createApplicationInSpaceArgsForCall, struct {
		app       v3action.Application
		spaceGUID string
	}{app, spaceGUID})
	fake.recordInvocation("CreateApplicationInSpace", []interface{}{app, spaceGUID})
	fake.createApplicationInSpaceMutex.Unlock()
	if fake.CreateApplicationInSpaceStub != nil {
		return fake.CreateApplicationInSpaceStub(app, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createApplicationInSpaceReturns.result1, fake.createApplicationInSpaceReturns.result2, fake.createApplicationInSpaceReturns.result3
}

func (fake *FakeV3PushActor) CreateApplicationInSpaceCallCount() int {
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	return len(fake.createApplicationInSpaceArgsForCall)
}

func (fake *FakeV3PushActor) CreateApplicationInSpaceArgsForCall(i int) (v3action.Application, string) {
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	return fake.createApplicationInSpaceArgsForCall[i].app, fake.createApplicationInSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV3PushActor) CreateApplicationInSpaceReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.CreateApplicationInSpaceStub = nil
	fake.createApplicationInSpaceReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) CreateApplicationInSpaceReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.CreateApplicationInSpaceStub = nil
	if fake.createApplicationInSpaceReturnsOnCall == nil {
		fake.createApplicationInSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.createApplicationInSpaceReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) GetApplicationByNameAndSpace(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
	}{appName, spaceGUID})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{appName, spaceGUID})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(appName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
}

func (fake *FakeV3PushActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3PushActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].appName, fake.getApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV3PushActor) GetApplicationByNameAndSpaceReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) GetApplicationSummaryByNameAndSpace(appName string, spaceGUID string) (v3action.ApplicationSummary, v3action.Warnings, error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)]
	fake.getApplicationSummaryByNameAndSpaceArgsForCall = append(fake.getApplicationSummaryByNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
	}{appName, spaceGUID})
	fake.recordInvocation("GetApplicationSummaryByNameAndSpace", []interface{}{appName, spaceGUID})
	fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationSummaryByNameAndSpaceStub != nil {
		return fake.GetApplicationSummaryByNameAndSpaceStub(appName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationSummaryByNameAndSpaceReturns.result1, fake.getApplicationSummaryByNameAndSpaceReturns.result2, fake.getApplicationSummaryByNameAndSpaceReturns.result3
}

func (fake *FakeV3PushActor) GetApplicationSummaryByNameAndSpaceCallCount() int {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)
}

func (fake *FakeV3PushActor) GetApplicationSummaryByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].appName, fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV3PushActor) GetApplicationSummaryByNameAndSpaceReturns(result1 v3action.ApplicationSummary, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	fake.getApplicationSummaryByNameAndSpaceReturns = struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) GetApplicationSummaryByNameAndSpaceReturnsOnCall(i int, result1 v3action.ApplicationSummary, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	if fake.getApplicationSummaryByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationSummaryByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.ApplicationSummary
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) GetStreamingLogsForApplicationByNameAndSpace(appName string, spaceGUID string, client v3action.NOAAClient) (<-chan *v3action.LogMessage, <-chan error, v3action.Warnings, error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)]
	fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall = append(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
		client    v3action.NOAAClient
	}{appName, spaceGUID, client})
	fake.recordInvocation("GetStreamingLogsForApplicationByNameAndSpace", []interface{}{appName, spaceGUID, client})
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetStreamingLogsForApplicationByNameAndSpaceStub != nil {
		return fake.GetStreamingLogsForApplicationByNameAndSpaceStub(appName, spaceGUID, client)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result1, fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result2, fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result3, fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result4
}

func (fake *FakeV3PushActor) GetStreamingLogsForApplicationByNameAndSpaceCallCount() int {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3PushActor) GetStreamingLogsForApplicationByNameAndSpaceArgsForCall(i int) (string, string, v3action.NOAAClient) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i].appName, fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i].spaceGUID, fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i].client
}

func (fake *FakeV3PushActor) GetStreamingLogsForApplicationByNameAndSpaceReturns(result1 <-chan *v3action.LogMessage, result2 <-chan error, result3 v3action.Warnings, result4 error) {
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	fake.getStreamingLogsForApplicationByNameAndSpaceReturns = struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeV3PushActor) GetStreamingLogsForApplicationByNameAndSpaceReturnsOnCall(i int, result1 <-chan *v3action.LogMessage, result2 <-chan error, result3 v3action.Warnings, result4 error) {
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	if fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 <-chan *v3action.LogMessage
			result2 <-chan error
			result3 v3action.Warnings
			result4 error
		})
	}
	fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeV3PushActor) PollStart(appGUID string, warnings chan<- v3action.Warnings) error {
	fake.pollStartMutex.Lock()
	ret, specificReturn := fake.pollStartReturnsOnCall[len(fake.pollStartArgsForCall)]
	fake.pollStartArgsForCall = append(fake.pollStartArgsForCall, struct {
		appGUID  string
		warnings chan<- v3action.Warnings
	}{appGUID, warnings})
	fake.recordInvocation("PollStart", []interface{}{appGUID, warnings})
	fake.pollStartMutex.Unlock()
	if fake.PollStartStub != nil {
		return fake.PollStartStub(appGUID, warnings)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pollStartReturns.result1
}

func (fake *FakeV3PushActor) PollStartCallCount() int {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	return len(fake.pollStartArgsForCall)
}

func (fake *FakeV3PushActor) PollStartArgsForCall(i int) (string, chan<- v3action.Warnings) {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	return fake.pollStartArgsForCall[i].appGUID, fake.pollStartArgsForCall[i].warnings
}

func (fake *FakeV3PushActor) PollStartReturns(result1 error) {
	fake.PollStartStub = nil
	fake.pollStartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeV3PushActor) PollStartReturnsOnCall(i int, result1 error) {
	fake.PollStartStub = nil
	if fake.pollStartReturnsOnCall == nil {
		fake.pollStartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pollStartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeV3PushActor) SetApplicationDroplet(appName string, spaceGUID string, dropletGUID string) (v3action.Warnings, error) {
	fake.setApplicationDropletMutex.Lock()
	ret, specificReturn := fake.setApplicationDropletReturnsOnCall[len(fake.setApplicationDropletArgsForCall)]
	fake.setApplicationDropletArgsForCall = append(fake.setApplicationDropletArgsForCall, struct {
		appName     string
		spaceGUID   string
		dropletGUID string
	}{appName, spaceGUID, dropletGUID})
	fake.recordInvocation("SetApplicationDroplet", []interface{}{appName, spaceGUID, dropletGUID})
	fake.setApplicationDropletMutex.Unlock()
	if fake.SetApplicationDropletStub != nil {
		return fake.SetApplicationDropletStub(appName, spaceGUID, dropletGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setApplicationDropletReturns.result1, fake.setApplicationDropletReturns.result2
}

func (fake *FakeV3PushActor) SetApplicationDropletCallCount() int {
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	return len(fake.setApplicationDropletArgsForCall)
}

func (fake *FakeV3PushActor) SetApplicationDropletArgsForCall(i int) (string, string, string) {
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	return fake.setApplicationDropletArgsForCall[i].appName, fake.setApplicationDropletArgsForCall[i].spaceGUID, fake.setApplicationDropletArgsForCall[i].dropletGUID
}

func (fake *FakeV3PushActor) SetApplicationDropletReturns(result1 v3action.Warnings, result2 error) {
	fake.SetApplicationDropletStub = nil
	fake.setApplicationDropletReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3PushActor) SetApplicationDropletReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.SetApplicationDropletStub = nil
	if fake.setApplicationDropletReturnsOnCall == nil {
		fake.setApplicationDropletReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.setApplicationDropletReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3PushActor) StagePackage(packageGUID string, appName string) (<-chan v3action.Droplet, <-chan v3action.Warnings, <-chan error) {
	fake.stagePackageMutex.Lock()
	ret, specificReturn := fake.stagePackageReturnsOnCall[len(fake.stagePackageArgsForCall)]
	fake.stagePackageArgsForCall = append(fake.stagePackageArgsForCall, struct {
		packageGUID string
		appName     string
	}{packageGUID, appName})
	fake.recordInvocation("StagePackage", []interface{}{packageGUID, appName})
	fake.stagePackageMutex.Unlock()
	if fake.StagePackageStub != nil {
		return fake.StagePackageStub(packageGUID, appName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.stagePackageReturns.result1, fake.stagePackageReturns.result2, fake.stagePackageReturns.result3
}

func (fake *FakeV3PushActor) StagePackageCallCount() int {
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	return len(fake.stagePackageArgsForCall)
}

func (fake *FakeV3PushActor) StagePackageArgsForCall(i int) (string, string) {
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	return fake.stagePackageArgsForCall[i].packageGUID, fake.stagePackageArgsForCall[i].appName
}

func (fake *FakeV3PushActor) StagePackageReturns(result1 <-chan v3action.Droplet, result2 <-chan v3action.Warnings, result3 <-chan error) {
	fake.StagePackageStub = nil
	fake.stagePackageReturns = struct {
		result1 <-chan v3action.Droplet
		result2 <-chan v3action.Warnings
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) StagePackageReturnsOnCall(i int, result1 <-chan v3action.Droplet, result2 <-chan v3action.Warnings, result3 <-chan error) {
	fake.StagePackageStub = nil
	if fake.stagePackageReturnsOnCall == nil {
		fake.stagePackageReturnsOnCall = make(map[int]struct {
			result1 <-chan v3action.Droplet
			result2 <-chan v3action.Warnings
			result3 <-chan error
		})
	}
	fake.stagePackageReturnsOnCall[i] = struct {
		result1 <-chan v3action.Droplet
		result2 <-chan v3action.Warnings
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) StartApplication(appGUID string) (v3action.Application, v3action.Warnings, error) {
	fake.startApplicationMutex.Lock()
	ret, specificReturn := fake.startApplicationReturnsOnCall[len(fake.startApplicationArgsForCall)]
	fake.startApplicationArgsForCall = append(fake.startApplicationArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("StartApplication", []interface{}{appGUID})
	fake.startApplicationMutex.Unlock()
	if fake.StartApplicationStub != nil {
		return fake.StartApplicationStub(appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.startApplicationReturns.result1, fake.startApplicationReturns.result2, fake.startApplicationReturns.result3
}

func (fake *FakeV3PushActor) StartApplicationCallCount() int {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	return len(fake.startApplicationArgsForCall)
}

func (fake *FakeV3PushActor) StartApplicationArgsForCall(i int) string {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	return fake.startApplicationArgsForCall[i].appGUID
}

func (fake *FakeV3PushActor) StartApplicationReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.StartApplicationStub = nil
	fake.startApplicationReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) StartApplicationReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.StartApplicationStub = nil
	if fake.startApplicationReturnsOnCall == nil {
		fake.startApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.startApplicationReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) StopApplication(appGUID string) (v3action.Warnings, error) {
	fake.stopApplicationMutex.Lock()
	ret, specificReturn := fake.stopApplicationReturnsOnCall[len(fake.stopApplicationArgsForCall)]
	fake.stopApplicationArgsForCall = append(fake.stopApplicationArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("StopApplication", []interface{}{appGUID})
	fake.stopApplicationMutex.Unlock()
	if fake.StopApplicationStub != nil {
		return fake.StopApplicationStub(appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.stopApplicationReturns.result1, fake.stopApplicationReturns.result2
}

func (fake *FakeV3PushActor) StopApplicationCallCount() int {
	fake.stopApplicationMutex.RLock()
	defer fake.stopApplicationMutex.RUnlock()
	return len(fake.stopApplicationArgsForCall)
}

func (fake *FakeV3PushActor) StopApplicationArgsForCall(i int) string {
	fake.stopApplicationMutex.RLock()
	defer fake.stopApplicationMutex.RUnlock()
	return fake.stopApplicationArgsForCall[i].appGUID
}

func (fake *FakeV3PushActor) StopApplicationReturns(result1 v3action.Warnings, result2 error) {
	fake.StopApplicationStub = nil
	fake.stopApplicationReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3PushActor) StopApplicationReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.StopApplicationStub = nil
	if fake.stopApplicationReturnsOnCall == nil {
		fake.stopApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.stopApplicationReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3PushActor) UpdateApplication(app v3action.Application) (v3action.Application, v3action.Warnings, error) {
	fake.updateApplicationMutex.Lock()
	ret, specificReturn := fake.updateApplicationReturnsOnCall[len(fake.updateApplicationArgsForCall)]
	fake.updateApplicationArgsForCall = append(fake.updateApplicationArgsForCall, struct {
		app v3action.Application
	}{app})
	fake.recordInvocation("UpdateApplication", []interface{}{app})
	fake.updateApplicationMutex.Unlock()
	if fake.UpdateApplicationStub != nil {
		return fake.UpdateApplicationStub(app)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.updateApplicationReturns.result1, fake.updateApplicationReturns.result2, fake.updateApplicationReturns.result3
}

func (fake *FakeV3PushActor) UpdateApplicationCallCount() int {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return len(fake.updateApplicationArgsForCall)
}

func (fake *FakeV3PushActor) UpdateApplicationArgsForCall(i int) v3action.Application {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return fake.updateApplicationArgsForCall[i].app
}

func (fake *FakeV3PushActor) UpdateApplicationReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.UpdateApplicationStub = nil
	fake.updateApplicationReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) UpdateApplicationReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.UpdateApplicationStub = nil
	if fake.updateApplicationReturnsOnCall == nil {
		fake.updateApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.updateApplicationReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3PushActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.createPackageByApplicationNameAndSpaceMutex.RLock()
	defer fake.createPackageByApplicationNameAndSpaceMutex.RUnlock()
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	fake.stopApplicationMutex.RLock()
	defer fake.stopApplicationMutex.RUnlock()
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3PushActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.V3PushActor = new(FakeV3PushActor)