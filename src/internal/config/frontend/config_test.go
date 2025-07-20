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

package frontend

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewConfig(t *testing.T) {
	testCases := map[string]struct {
		in   int
		want *Config
	}{
		"zero": {0, newDefaultWantConfig()},
		"one":  {1, newDefaultWantConfig()},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := NewConfig(testCase.in)
			if diff := cmp.Diff(testCase.want, got); diff != "" {
				t.Errorf("NewConfig(%d) = diff -want +got\n%s", testCase.in, diff)
			}
		})
	}
}

func newDefaultWantConfig() *Config {
	return &Config{
		UI:        newDefaultWantUI(),
		Shortcuts: nil,
		Valid:     true,
	}
}
