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

	"github.com/jdheim/launchee/internal/config/yaml"
)

func TestConfigPathStub(t *testing.T) {
	testCases := map[string]struct {
		input yaml.ConfigPath
		want  []string
	}{
		"valid":             {ConfigPathValidStub{}, []string{launchee7ConfigPath, launcheeOverrideConfigPath}},
		"invalid":           {ConfigPathInvalidStub{}, []string{invalidConfigPath, invalidConfigPath}},
		"system invalid":    {SystemConfigPathInvalidStub{}, []string{invalidConfigPath, launchee7ConfigPath}},
		"user invalid":      {UserConfigPathInvalidStub{}, []string{launchee7ConfigPath, invalidConfigPath}},
		"not exists":        {ConfigPathNotExistsStub{}, []string{notExistsConfigPath, notExistsConfigPath}},
		"system not exists": {SystemConfigPathNotExistsStub{}, []string{notExistsConfigPath, launchee7ConfigPath}},
		"user not exists":   {UserConfigPathNotExistsStub{}, []string{launchee7ConfigPath, notExistsConfigPath}},
		"empty":             {ConfigPathEmptyStub{}, []string{launcheeEmptyConfigPath, launcheeEmptyConfigPath}},
		"system empty":      {SystemConfigPathEmptyStub{}, []string{launcheeEmptyConfigPath, launchee7ConfigPath}},
		"user empty":        {UserConfigPathEmptyStub{}, []string{launchee7ConfigPath, launcheeEmptyConfigPath}},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			systemConfigPath := testCase.input.GetSystemConfigPath()
			if systemConfigPath != testCase.want[0] {
				t.Errorf("GetSystemConfigPath() = %q, want %q", systemConfigPath, testCase.want[0])
			}
			userConfigPath := testCase.input.GetUserConfigPath()
			if userConfigPath != testCase.want[1] {
				t.Errorf("GetUserConfigPath() = %q, want %q", systemConfigPath, testCase.want[1])
			}
		})
	}
}
