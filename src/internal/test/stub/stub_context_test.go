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

package stub

import (
	"testing"
)

func TestContextStub(t *testing.T) {
	testCases := []string{"logger", "frontend"}

	for _, name := range testCases {
		t.Run(name, func(t *testing.T) {
			got := ContextStub{}.New().Value(name)
			want := struct{}{}
			if got != want {
				t.Errorf("New() = %v, want %v", got, want)
			}
		})
	}
}
