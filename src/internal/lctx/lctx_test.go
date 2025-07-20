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
	"errors"
	"fmt"
	"testing"

	"github.com/jdheim/launchee/internal/test/assert"
	"github.com/jdheim/launchee/internal/test/stub"
)

func TestGetContext(t *testing.T) {
	want := context.WithValue(context.Background(), "k1", "v1")
	SetContext(want)
	got := GetContext()

	if got != want {
		t.Error("GetContext() = equal expected")
	}
	if got == context.WithValue(context.Background(), "k2", "v2") {
		t.Error("GetContext() = not equal expected")
	}
}

func TestLogging(t *testing.T) {
	testCases := map[string]struct {
		input func()
	}{
		"LogInfo": {func() {
			LogInfo("test message")
		}},
		"LogInfof": {func() {
			LogInfof("test %s", "message")
		}},
		"LogError": {func() {
			LogError("test message")
		}},
		"LogErrorf": {func() {
			LogErrorf("test %s", "message")
		}},
	}

	SetContext(stub.ContextStub{}.New())
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			defer assert.FuncPanic(t, fmt.Sprintf("%s()", name), fmt.Sprintf("runtime.%s", name))
			testCase.input()
		})
	}
}

func TestNewErrorMessageDialog(t *testing.T) {
	testCases := []string{"valid", "invalid"}

	SetContext(stub.ContextStub{}.New())
	for _, name := range testCases {
		t.Run(name, func(t *testing.T) {
			if name == "valid" {
				defer assert.FuncPanic(t, "NewErrorMessageDialog()", "runtime.MessageDialog")
			} else if name == "invalid" {
				MessageDialogImpl = stub.MessageDialogErrorStub{}
				LoggerImpl = stub.LoggerStub{}
			}
			NewErrorMessageDialog("Error occurred during application startup", errors.ErrUnsupported)
		})
	}
}
