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
	"fmt"
	"os"
	"os/exec"
	"testing"
)

type TestedFunc struct {
	FunctionName string
	Function     func()
	ExitCode     int
}

func FuncExited(t *testing.T, testedFunc *TestedFunc) error {
	t.Helper()
	if os.Getenv(t.Name()) == "1" {
		testedFunc.Function()
		return nil
	}
	command := exec.Command(os.Args[0], fmt.Sprintf("-test.run=^%s$", t.Name()))
	command.Env = append(os.Environ(), fmt.Sprintf("%s=1", t.Name()))
	_ = command.Run()
	want := testedFunc.ExitCode
	got := command.ProcessState.ExitCode()
	if got != want {
		return fmt.Errorf("%s() = %d, want exit status %d", testedFunc.FunctionName, got, want)
	}
	return nil
}
