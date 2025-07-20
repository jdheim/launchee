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

type WindowStub struct{}

func (WindowStub) SetTitle(title string) {
	if debug.IsDebugEnabled() {
		log.Printf("SetTitle: %s", title)
	}
}

func (WindowStub) SetSize(width int, height int) {
	if debug.IsDebugEnabled() {
		log.Printf("SetSize: width=%d, height=%d", width, height)
	}
}

func (WindowStub) SetMinSize(width int, height int) {
	if debug.IsDebugEnabled() {
		log.Printf("SetMinSize: width=%d, height=%d", width, height)
	}
}

func (WindowStub) SetMaxSize(width int, height int) {
	if debug.IsDebugEnabled() {
		log.Printf("SetMaxSize: width=%d, height=%d", width, height)
	}
}

func (WindowStub) Quit() {
	if debug.IsDebugEnabled() {
		log.Printf("Quiting")
	}
}
