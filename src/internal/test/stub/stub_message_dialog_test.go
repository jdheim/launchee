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

	"github.com/jdheim/launchee/internal/lctx"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func TestMessageDialogStub(t *testing.T) {
	testCases := map[string]struct {
		input lctx.MessageDialog
	}{
		"valid": {MessageDialogValidStub{}},
		"error": {MessageDialogErrorStub{}},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := testCase.input.Open(ContextStub{}.New(), runtime.MessageDialogOptions{})
			want := "ok"
			if "valid" == name {
				if err != nil {
					t.Errorf("Open() = %v", err)
				} else if result != want {
					t.Errorf("Open() = %q, want %q", result, want)
				}
			} else if "error" == name && err == nil {
				t.Error("Open() = error expected")
			}
		})
	}
}
