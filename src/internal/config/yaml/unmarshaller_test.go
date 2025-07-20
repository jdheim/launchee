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

package yaml

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jdheim/launchee/internal/lctx"
	"github.com/jdheim/launchee/internal/test/stub"
)

func TestUnmarshalCustomConfig(t *testing.T) {
	defer chdirBack(t)
	chdirToRoot(t)
	testCases := map[string]struct {
		configPathStub ConfigPath
		wantErr        bool
	}{
		"valid":      {stub.ConfigPathValidStub{}, false},
		"invalid":    {stub.ConfigPathInvalidStub{}, true},
		"not exists": {stub.ConfigPathNotExistsStub{}, false},
		"empty":      {stub.ConfigPathEmptyStub{}, false},
	}

	lctx.LoggerImpl = stub.LoggerStub{}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			_, err := UnmarshalCustomConfig(testCase.configPathStub.GetSystemConfigPath())
			if gotErr := err == nil; gotErr == testCase.wantErr {
				t.Errorf("UnmarshalCustomConfig() = %t, want %t", gotErr, testCase.wantErr)
			}
		})
	}
}

func TestUnmarshalConfigs(t *testing.T) {
	defer chdirBack(t)
	chdirToRoot(t)
	testCases := map[string]struct {
		configPathStub ConfigPath
		wantErr        bool
	}{
		"valid":             {stub.ConfigPathValidStub{}, false},
		"invalid":           {stub.ConfigPathInvalidStub{}, true},
		"system invalid":    {stub.SystemConfigPathInvalidStub{}, true},
		"user invalid":      {stub.UserConfigPathInvalidStub{}, true},
		"not exists":        {stub.ConfigPathNotExistsStub{}, false},
		"system not exists": {stub.SystemConfigPathNotExistsStub{}, false},
		"user not exists":   {stub.UserConfigPathNotExistsStub{}, false},
		"empty":             {stub.ConfigPathEmptyStub{}, false},
		"system empty":      {stub.SystemConfigPathEmptyStub{}, false},
		"user empty":        {stub.UserConfigPathEmptyStub{}, false},
	}

	lctx.LoggerImpl = stub.LoggerStub{}
	defer func() { ConfigPathImpl = systemAwareConfigPath{} }()
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ConfigPathImpl = testCase.configPathStub
			_, err := UnmarshalConfigs()
			if gotErr := err == nil; gotErr == testCase.wantErr {
				t.Errorf("UnmarshalConfigs() = %t, want %t", gotErr, testCase.wantErr)
			}
		})
	}
}

func chdirBack(t *testing.T) {
	if err := os.Chdir(filepath.Join("internal", "config", "yaml")); err != nil {
		t.Errorf("UnmarshalConfigs() = %v", err)
	}
}

func chdirToRoot(t *testing.T) {
	if err := os.Chdir(filepath.Join("..", "..", "..")); err != nil {
		t.Errorf("UnmarshalConfigs() = %v", err)
	}
}

func TestGetGOOS(t *testing.T) {
	got := getGOOS()
	want := "linux"
	if got != want {
		t.Errorf("getGOOS() = %q, want %q", got, want)
	}
}

func TestSystemAwareConfigPath(t *testing.T) {
	testCases := []string{"linux", "windows", "darwin"}

	for _, name := range testCases {
		t.Run(name, func(t *testing.T) {
			getGOOS = func() string {
				return name
			}
			if getGOOS() == "windows" {
				defer func() {
					err := os.Unsetenv("PROGRAMDATA")
					if err != nil {
						t.Errorf("Unsetenv() = %v", err)
					}
				}()
				err := os.Setenv("PROGRAMDATA", "C:\\ProgramData")
				if err != nil {
					t.Errorf("Setenv() = %v", err)
				}
			}
			ConfigPathImpl.GetSystemConfigPath()
			ConfigPathImpl.GetUserConfigPath()
		})
	}
}

func TestFindConfigFile(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  string
	}{
		"empty":      {"", ""},
		"not-exists": {t.TempDir(), ""},
		"yml":        {t.TempDir(), "launchee/launchee.yml"},
		"yaml":       {t.TempDir(), "launchee/launchee.yaml"},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			var want string
			if testCase.want != "" {
				configDir := filepath.Join(testCase.input, "launchee")
				if err := os.Mkdir(configDir, 0755); err != nil {
					t.Errorf("Mkdir() = %v", err)
				}
				want = filepath.Join(testCase.input, testCase.want)
				if err := os.WriteFile(want, []byte{}, 0644); err != nil {
					t.Errorf("WriteFile() = %v", err)
				}
			}
			got := findConfigFile(testCase.input)
			if got != want {
				t.Errorf("findConfigFile(%q) = %q, want %q", testCase.input, got, want)
			}
		})
	}
}
