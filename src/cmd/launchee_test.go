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

package cmd

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jdheim/launchee/internal/config/frontend"
	"github.com/jdheim/launchee/internal/config/yaml"
	"github.com/jdheim/launchee/internal/lctx"
	"github.com/jdheim/launchee/internal/test/stub"
)

func TestNewLaunchee(t *testing.T) {
	got := NewLaunchee()
	want := &Launchee{}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("NewLaunchee() = diff -want +got\n%s", diff)
	}
}

func TestStartup(t *testing.T) {
	testCases := []string{"valid", "invalid"}

	for _, name := range testCases {
		t.Run(name, func(t *testing.T) {
			windowImpl = stub.WindowStub{}
			lctx.LoggerImpl = stub.LoggerStub{}
			if name == "invalid" {
				lctx.MessageDialogImpl = stub.MessageDialogErrorStub{}
				yaml.ConfigPathImpl = stub.ConfigPathInvalidStub{}
			}
			NewLaunchee().Startup(stub.ContextStub{}.New())
		})
	}

	lctx.LoggerImpl = stub.LoggerStub{}
	windowImpl = stub.WindowStub{}
	NewLaunchee().Startup(stub.ContextStub{}.New())
}

func TestIsBuildForJdvm(t *testing.T) {
	buildForJdvm = "true"
	got := NewLaunchee().IsBuildForJdvm()
	want := true
	if got != want {
		t.Errorf("IsBuildForJdvm() = %t, want %t", got, true)
	}
}

func TestGetAppVersion(t *testing.T) {
	appVersion = "0.0.1"
	want := appVersion
	got := NewLaunchee().GetAppVersion()
	if got != want {
		t.Errorf("GetAppVersion() = %s, want %s", got, want)
	}
}

func TestGetConfig(t *testing.T) {
	want := &frontend.Config{UI: nil, Shortcuts: nil, Valid: true}
	testLaunchee := &Launchee{Config: &frontend.Config{UI: nil, Shortcuts: nil, Valid: true}}
	got := testLaunchee.GetConfig()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetConfig() = diff -want +got\n%s", diff)
	}
}

func TestRunCommand(t *testing.T) {
	testLaunchee := &Launchee{}
	testLaunchee.RunCommand("echo", []string{"-n", "test"})
}
