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

package debug

import (
	"os"
	"testing"
)

func TestIsDebugEnabled(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  bool
	}{
		"debug enabled":  {"DEBUG", true},
		"debug disabled": {"NO_DEBUG", false},
		"empty":          {"", false},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			defer func() {
				err := os.Unsetenv(testCase.input)
				if err != nil {
					t.Errorf("IsDebugEnabled() = %v", err)
				}
			}()
			if testCase.input != "" {
				err := os.Setenv(testCase.input, "")
				if err != nil {
					t.Errorf("IsDebugEnabled() = %v", err)
				}
			}
			got := IsDebugEnabled()
			if got != testCase.want {
				t.Errorf("IsDebugEnabled() = %t, want %t", got, testCase.want)
			}
		})
	}
}

func TestEnableDebug(t *testing.T) {
	EnableDebug()
	got := IsDebugEnabled()
	if got != true {
		t.Errorf("IsDebugEnabled() = %t, want %t", got, true)
	}
}
