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
	"runtime/debug"
	"strings"
	"testing"
)

func FuncPanic(t *testing.T, funcName string, want string) {
	t.Helper()
	stackTrace := string(debug.Stack())
	if recovered := recover(); recovered == nil {
		t.Errorf("%s = panic expected", funcName)
	}
	if !strings.Contains(stackTrace, want) {
		t.Errorf("%s = %q not in the stack:\n%s", funcName, want, stackTrace)
	}
}
