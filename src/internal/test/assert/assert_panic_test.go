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

package assert

import (
	"testing"
)

func TestFuncPanic(t *testing.T) {
	testCases := map[string]struct {
		inputFunc     func()
		inputFuncName string
		want          bool
	}{
		"valid":             {func() { panic("test panic") }, "assert.FuncPanic", true},
		"invalid func name": {func() { panic("test panic") }, "assert.FuncPanicNotExists", false},
		"invalid no panic":  {func() {}, "assert.FuncPanic", false},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := FuncPanic(t, "TestFuncPanic()", testCase.inputFuncName, testCase.inputFunc)
			if got := err == nil; got != testCase.want {
				t.Errorf("FuncPanic() = %t, want %t", got, testCase.want)
			}
		})
	}
}
