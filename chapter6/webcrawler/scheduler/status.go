package scheduler

import (
	"fmt"
	"sync"
)

// Status 代表调度器状态的类型。
type Status uint8

const (
	// SchedStatusUninitialized 代表未初始化的状态。
	SchedStatusUninitialized Status = 0
	// SchedStatusInitializing 代表正在初始化的状态。
	SchedStatusInitializing Status = 1
	// SchedStatusInitialized 代表已初始化的状态。
	SchedStatusInitialized Status = 2
	// SchedStatusStarting 代表正在启动的状态。
	SchedStatusStarting Status = 3
	// SchedStatusStarted 代表已启动的状态。
	SchedStatusStarted Status = 4
	// SchedStatusStopping 代表正在停止的状态。
	SchedStatusStopping Status = 5
	// SchedStatusStopped 代表已停止的状态。
	SchedStatusStopped Status = 6
)

// checkStatus 用于状态的检查。
// 参数currentStatus代表当前的状态。
// 参数wantedStatus代表想要的状态。
// 检查规则：
//     1. 处于正在初始化、正在启动或正在停止状态时，不能从外部改变状态。
//     2. 想要的状态只能是正在初始化、正在启动或正在停止状态中的一个。
//     3. 处于未初始化状态时，不能变为正在启动或正在停止状态。
//     4. 处于已启动状态时，不能变为正在初始化或正在启动状态。
//     5. 只要未处于已启动状态就不能变为正在停止状态。
func checkStatus(
	currentStatus Status,
	wantedStatus Status,
	lock sync.Locker) (err error) {
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	switch currentStatus {
	case SchedStatusInitializing:
		err = genError("the scheduler is being initialized!")
	case SchedStatusStarting:
		err = genError("the scheduler is being started!")
	case SchedStatusStopping:
		err = genError("the scheduler is being stopped!")
	}
	if err != nil {
		return
	}
	if currentStatus == SchedStatusUninitialized &&
		(wantedStatus == SchedStatusStarting ||
			wantedStatus == SchedStatusStopping) {
		err = genError("the scheduler has not yet been initialized!")
		return
	}
	switch wantedStatus {
	case SchedStatusInitializing:
		switch currentStatus {
		case SchedStatusStarted:
			err = genError("the scheduler has been started!")
		}
	case SchedStatusStarting:
		switch currentStatus {
		case SchedStatusUninitialized:
			err = genError("the scheduler has not been initialized!")
		case SchedStatusStarted:
			err = genError("the scheduler has been started!")
		}
	case SchedStatusStopping:
		if currentStatus != SchedStatusStarted {
			err = genError("the scheduler has not been started!")
		}
	default:
		errMsg :=
			fmt.Sprintf("unsupported wanted status for check! (wantedStatus: %d)",
				wantedStatus)
		err = genError(errMsg)
	}
	return
}

// GetStatusDescription 用于获取状态的文字描述。
func GetStatusDescription(status Status) string {
	switch status {
	case SchedStatusUninitialized:
		return "uninitialized"
	case SchedStatusInitializing:
		return "initializing"
	case SchedStatusInitialized:
		return "initialized"
	case SchedStatusStarting:
		return "starting"
	case SchedStatusStarted:
		return "started"
	case SchedStatusStopping:
		return "stopping"
	case SchedStatusStopped:
		return "stopped"
	default:
		return "unknown"
	}
}
