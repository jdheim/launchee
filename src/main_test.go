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

package main

import (
	"errors"
	"testing"

	"github.com/jdheim/launchee/cmd"
	"github.com/jdheim/launchee/internal/config/frontend"
	"github.com/jdheim/launchee/internal/test/assert"
	"github.com/jdheim/launchee/internal/test/stub"
	flag "github.com/spf13/pflag"
)

func TestWailsMainStartGUI(t *testing.T) {
	if err := assert.FuncExited(t, &assert.TestedFunc{
		FunctionName: "main",
		Function:     func() { main() },
		ExitCode:     1,
	}); err != nil {
		t.Error(err)
	}
}

func TestProcessCommandLineFlags(t *testing.T) {
	testCases := map[string]struct {
		input    []string
		exitCode int
	}{
		"help short":              {[]string{"-h"}, 0},
		"help long":               {[]string{"--help"}, 0},
		"incorrect help short":    {[]string{"--h"}, 2},
		"incorrect help long":     {[]string{"-help"}, 2},
		"version short":           {[]string{"-v"}, 0},
		"version long":            {[]string{"--version"}, 0},
		"incorrect version short": {[]string{"--v"}, 2},
		"incorrect version long":  {[]string{"-version"}, 2},
		"config short":            {[]string{"-c", stub.ConfigPathValidStub{}.GetSystemConfigPath()}, -1},
		"config long":             {[]string{"--config", stub.ConfigPathValidStub{}.GetSystemConfigPath()}, -1},
		"invalid config short":    {[]string{"-c", stub.ConfigPathNotExistsStub{}.GetSystemConfigPath()}, 1},
		"invalid config long":     {[]string{"--config", stub.ConfigPathNotExistsStub{}.GetSystemConfigPath()}, 1},
		"incorrect flag short":    {[]string{"-e"}, 2},
		"incorrect flag long":     {[]string{"--error"}, 2},
	}

	t.Cleanup(func() { parse = flag.Parse })
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet(t.Name(), flag.ExitOnError)
			if testCase.exitCode == -1 {
				parse = testParse(t, testCase.input)
				processCommandLineFlags()
			} else {
				if err := assert.FuncExited(t, &assert.TestedFunc{
					FunctionName: "processCommandLineFlags",
					Function: func() {
						parse = testParse(t, testCase.input)
						processCommandLineFlags()
					},
					ExitCode: testCase.exitCode,
				}); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

func testParse(t *testing.T, input []string) func() {
	t.Helper()
	return func() {
		if err := flag.CommandLine.Parse(input); err != nil {
			t.Error(err)
		}
	}
}

func TestAssertLauncheeConfigValid(t *testing.T) {
	if err := assert.FuncExited(t, &assert.TestedFunc{
		FunctionName: "assertLauncheeConfigValid",
		Function: func() {
			assertLauncheeConfigValid(&cmd.Launchee{
				Config: &frontend.Config{
					Valid: false,
				},
			})
		},
		ExitCode: 1,
	}); err != nil {
		t.Error(err)
	}
}

func TestAssertError(t *testing.T) {
	if err := assert.FuncExited(t, &assert.TestedFunc{
		FunctionName: "assertError",
		Function: func() {
			assertError(errors.ErrUnsupported)
		},
		ExitCode: 1,
	}); err != nil {
		t.Error(err)
	}
}
