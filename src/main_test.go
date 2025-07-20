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
)

func TestWailsMain(t *testing.T) {
	assert.FuncExited(t, &assert.TestedFunc{
		TestName:     "TestWailsMain",
		FunctionName: "main",
		Function: func() {
			main()
		},
		ExitCode: 1,
	})
}

func TestAssertLauncheeConfigValid(t *testing.T) {
	assert.FuncExited(t, &assert.TestedFunc{
		TestName:     "TestAssertLauncheeConfigValid",
		FunctionName: "assertLauncheeConfigValid",
		Function: func() {
			assertLauncheeConfigValid(&cmd.Launchee{
				Config: &frontend.Config{
					Valid: false,
				},
			})
		},
		ExitCode: 1,
	})
}

func TestAssertError(t *testing.T) {
	assert.FuncExited(t, &assert.TestedFunc{
		TestName:     "TestAssertError",
		FunctionName: "assertError",
		Function: func() {
			assertError(errors.ErrUnsupported)
		},
		ExitCode: 1,
	})
}
