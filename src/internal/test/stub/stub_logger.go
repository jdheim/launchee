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
	"log"

	"github.com/jdheim/launchee/internal/test/debug"
)

type LoggerStub struct{}

func (LoggerStub) LogInfo(message string) {
	if debug.IsDebugEnabled() {
		log.Print(message)
	}
}

func (LoggerStub) LogInfof(message string, args ...interface{}) {
	if debug.IsDebugEnabled() {
		log.Printf(message, args...)
	}
}

func (LoggerStub) LogError(message string) {
	if debug.IsDebugEnabled() {
		log.Print("ERR | " + message)
	}
}

func (LoggerStub) LogErrorf(message string, args ...interface{}) {
	if debug.IsDebugEnabled() {
		log.Printf("ERR | "+message, args...)
	}
}
