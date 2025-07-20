/*
 * © 2025-2025 JDHeim.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lctx

import (
	"context"
	"fmt"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type launcheeContext struct {
	ctx context.Context
}

var lock = &sync.Mutex{}
var singleInstance *launcheeContext

// getInstance returns the singleton instance of the launcheeContext.
func getInstance() *launcheeContext {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &launcheeContext{}
		}
	}

	return singleInstance
}

// GetContext returns the context.Context.
func GetContext() context.Context {
	return getInstance().ctx
}

// SetContext sets the context.Context.
func SetContext(ctx context.Context) {
	getInstance().ctx = ctx
}

type Logger interface {
	LogInfo(message string)
	LogInfof(format string, args ...interface{})
	LogError(message string)
	LogErrorf(format string, args ...interface{})
}

type runtimeLogger struct{}

func (runtimeLogger) LogInfo(message string) {
	runtime.LogInfo(GetContext(), message)
}

func (runtimeLogger) LogInfof(message string, args ...interface{}) {
	runtime.LogInfof(GetContext(), message, args...)
}

func (runtimeLogger) LogError(message string) {
	runtime.LogError(GetContext(), message)
}

func (runtimeLogger) LogErrorf(message string, args ...interface{}) {
	runtime.LogErrorf(GetContext(), message, args...)
}

var LoggerImpl Logger = runtimeLogger{}

func LogInfo(message string) {
	LoggerImpl.LogInfo(message)
}

func LogInfof(format string, args ...interface{}) {
	LoggerImpl.LogInfof(format, args...)
}

func LogError(message string) {
	LoggerImpl.LogError(message)
}

func LogErrorf(format string, args ...interface{}) {
	LoggerImpl.LogErrorf(format, args...)
}

type MessageDialog interface {
	Open(ctx context.Context, dialogOptions runtime.MessageDialogOptions) (string, error)
}

type runtimeMessageDialog struct{}

func (runtimeMessageDialog) Open(ctx context.Context, dialogOptions runtime.MessageDialogOptions) (string, error) {
	return runtime.MessageDialog(ctx, dialogOptions)
}

var MessageDialogImpl MessageDialog = runtimeMessageDialog{}

func NewErrorMessageDialog(message string, err error) {
	if _, err := MessageDialogImpl.Open(GetContext(), runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "Oops... Something Went Wrong",
		Message: fmt.Sprintf("%s:\n%v", message, err),
	}); err != nil {
		LogErrorf("Error while displaying message dialog: %v", err)
	}
}
